package main

import (
	"fmt"
	"net"
	"os"
	"runtime/debug"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/app/apiserver"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authentication"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authorization"
	errorwrap "github.com/nikita5637/quiz-registrator-api/internal/app/middleware/error_wrap"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/log"
	remindmanager "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager"
	game_reminder "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/game"
	lottery_reminder "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/lottery"
	adminservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/admin"
	certificatemanagerservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/certificate_manager"
	croupierservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/croupier"
	gameservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/game"
	gameplayerservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/game_player"
	gameresultmanagerservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/game_result_manager"
	leagueservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/league"
	photomanagerservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/photo_manager"
	placeservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/place"
	usermanagerservice "github.com/nikita5637/quiz-registrator-api/internal/app/service/user_manager"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/quiz_please"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/squiz"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/elasticsearch"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gamephotos"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameresults"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/userroles"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	rabbitmqproducer "github.com/nikita5637/quiz-registrator-api/internal/pkg/rabbitmq/producer"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/posener/ctxutil"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func init() {
	pflag.StringP("config", "c", "", "path to config file")
	_ = viper.BindPFlag("config", pflag.Lookup("config"))
}

func main() {
	ctx := ctxutil.Interrupt()

	pflag.Parse()

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := viper.GetBool("log.elastic.enabled")
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err := elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   viper.GetString("log.elastic.index"),
		})
		if err != nil {
			panic(err)
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", viper.GetString("log.module_name")),
	)))
	logger.InfoKV(ctx, "initialized logger", "log level", logLevel)

	driverName := viper.GetString("database.driver")
	db, err := storage.NewDB(ctx, driverName)
	if err != nil {
		logger.Fatalf(ctx, "new DB initialization error: %s", err.Error())
	}
	defer db.Close()

	rabbitMQConn, err := amqp.Dial(config.GetRabbitMQURL())
	if err != nil {
		logger.Fatalf(ctx, "get rabbitMQ connection error: %s", err.Error())
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		logger.Fatalf(ctx, "get rabbitMQ channel error: %s", err.Error())
	}
	defer rabbitMQChannel.Close()

	txManager := tx.NewManager(db)

	gameStorage := storage.NewGameStorage(driverName, txManager)
	gamePlayerStorage := storage.NewGamePlayerStorage(driverName, txManager)

	gamePlayersFacadeConfig := gameplayers.Config{
		GamePlayerStorage: gamePlayerStorage,
		TxManager:         txManager,
	}
	gamePlayersFacade := gameplayers.New(gamePlayersFacadeConfig)

	gamesFacadeConfig := games.Config{
		GameStorage: gameStorage,
		TxManager:   txManager,
	}
	gamesFacade := games.New(gamesFacadeConfig)

	quizPleaseCroupierConfig := quiz_please.Config{
		LotteryLink: quiz_please.LotteryLink,
	}

	squizCroupierConfig := squiz.Config{
		LotteryInfoPageLink:     squiz.LotteryInfoPageLink,
		LotteryRegistrationLink: squiz.LotteryRegistrationLink,
	}

	croupierConfig := croupier.Config{
		GamesFacade:        gamesFacade,
		QuizPleaseCroupier: quiz_please.New(quizPleaseCroupierConfig),
		SquizCroupier:      squiz.New(squizCroupierConfig),
	}
	croupier := croupier.New(croupierConfig)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		certificateStorage := storage.NewCertificateStorage(driverName, txManager)
		gamePhotoStorage := storage.NewGamePhotoStorage(driverName, txManager)
		gameResultStorage := storage.NewGameResultStorage(driverName, txManager)
		leagueStorage := storage.NewLeagueStorage(driverName, txManager)
		placeStorage := storage.NewPlaceStorage(driverName, txManager)
		userStorage := storage.NewUserStorage(driverName, txManager)
		userRoleStorage := storage.NewUserRoleStorage(driverName, txManager)

		certificatesFacadeConfig := certificates.Config{
			CertificateStorage: certificateStorage,
			TxManager:          txManager,
		}
		certificatesFacade := certificates.NewFacade(certificatesFacadeConfig)

		userRolesFacadeConfig := userroles.Config{
			TxManager:       txManager,
			UserRoleStorage: userRoleStorage,
		}
		userRolesFacade := userroles.New(userRolesFacadeConfig)

		adminServiceConfig := adminservice.Config{
			UserRolesFacade: userRolesFacade,
		}
		adminService := adminservice.New(adminServiceConfig)

		certificateManagerServiceConfig := certificatemanagerservice.Config{
			CertificatesFacade: certificatesFacade,
		}
		certificateManagerService := certificatemanagerservice.New(certificateManagerServiceConfig)

		croupierServiceConfig := croupierservice.Config{
			Croupier:          croupier,
			GamePlayersFacade: gamePlayersFacade,
			GamesFacade:       gamesFacade,
		}
		croupierService := croupierservice.New(croupierServiceConfig)

		gamePhotosFacadeConfig := gamephotos.Config{
			GameStorage:      gameStorage,
			GamePhotoStorage: gamePhotoStorage,
			TxManager:        txManager,
		}
		gamePhotosFacade := gamephotos.NewFacade(gamePhotosFacadeConfig)

		icsRabbitMQProducerConfig := rabbitmqproducer.Config{
			QueueName:       viper.GetString("service.game.ics.queue.name"),
			RabbitMQChannel: rabbitMQChannel,
		}
		icsRabbitMQProducer := rabbitmqproducer.New(icsRabbitMQProducerConfig)

		if err := icsRabbitMQProducer.Connect(ctx); err != nil {
			return fmt.Errorf("ICS producer connect error: %w", err)
		}

		gameServiceConfig := gameservice.Config{
			GamesFacade:      gamesFacade,
			RabbitMQProducer: icsRabbitMQProducer,
		}
		gameService := gameservice.New(gameServiceConfig)

		gamePlayerServiceConfig := gameplayerservice.Config{
			GamesFacade:       gamesFacade,
			GamePlayersFacade: gamePlayersFacade,
		}
		gamePlayerService := gameplayerservice.New(gamePlayerServiceConfig)

		gameResultsFacadeConfig := gameresults.Config{
			GameResultStorage: gameResultStorage,
			TxManager:         txManager,
		}
		gameResultsFacade := gameresults.NewFacade(gameResultsFacadeConfig)

		gameResultManagerServiceConfig := gameresultmanagerservice.Config{
			GameResultsFacade: gameResultsFacade,
		}
		gameResultManagerService := gameresultmanagerservice.New(gameResultManagerServiceConfig)

		leaguesFacadeConfig := leagues.Config{
			LeagueStorage: leagueStorage,
		}
		leaguesFacade := leagues.NewFacade(leaguesFacadeConfig)

		leagueServiceConfig := leagueservice.Config{
			LeaguesFacade: leaguesFacade,
		}
		leagueService := leagueservice.New(leagueServiceConfig)

		photoManagerServiceConfig := photomanagerservice.Config{
			GamePhotosFacade: gamePhotosFacade,
		}
		photoManagerService := photomanagerservice.New(photoManagerServiceConfig)

		placesFacadeConfig := places.Config{
			PlaceStorage: placeStorage,
		}
		placesFacade := places.NewFacade(placesFacadeConfig)

		placeServiceConfig := placeservice.Config{
			PlacesFacade: placesFacade,
		}
		placeService := placeservice.New(placeServiceConfig)

		usersFacadeConfig := users.Config{
			UserStorage: userStorage,
			TxManager:   txManager,
		}
		usersFacade := users.NewFacade(usersFacadeConfig)

		userManagerServiceConfig := usermanagerservice.Config{
			UsersFacade: usersFacade,
		}
		userManagerService := usermanagerservice.New(userManagerServiceConfig)

		authenticationMiddleware := authentication.New(authentication.Config{
			UsersFacade: usersFacade,
		})

		authorizationMiddleware := authorization.New(authorization.Config{
			UserRolesFacade: userRolesFacade,
		})

		errorWrapMiddleware := errorwrap.New()

		logMiddleware := log.New()

		apiServerConfig := apiserver.Config{
			AuthenticationMiddleware: authenticationMiddleware,
			AuthorizationMiddleware:  authorizationMiddleware,
			ErrorWrapMiddleware:      errorWrapMiddleware,
			LogMiddleware:            logMiddleware,

			AdminService:                 adminService,
			CertificateManagerService:    certificateManagerService,
			CroupierService:              croupierService,
			GameService:                  gameService,
			GameRegistratorService:       gameService,
			GamePlayerService:            gamePlayerService,
			GamePlayerRegistratorService: gamePlayerService,
			GameResultManagerService:     gameResultManagerService,
			LeagueService:                leagueService,
			PhotoManagerService:          photoManagerService,
			PlaceService:                 placeService,
			UserManagerService:           userManagerService,
		}
		apiServer := apiserver.New(apiServerConfig)

		lis, err := net.Listen("tcp", config.GetBindAddress())
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		logger.Infof(ctx, "starting registrator")
		return apiServer.ListenAndServe(ctx, lis)
	})

	g.Go(func() error {
		gameReminderRabbitMQProducerConfig := rabbitmqproducer.Config{
			QueueName:       viper.GetString("remind_manager.game.queue.name"),
			RabbitMQChannel: rabbitMQChannel,
		}
		gameReminderRabbitMQProducer := rabbitmqproducer.New(gameReminderRabbitMQProducerConfig)

		if err := gameReminderRabbitMQProducer.Connect(ctx); err != nil {
			return fmt.Errorf("game reminder producer connect error: %w", err)
		}

		gameReminderConfig := game_reminder.Config{
			GamesFacade:       gamesFacade,
			GamePlayersFacade: gamePlayersFacade,
			RabbitMQProducer:  gameReminderRabbitMQProducer,
		}
		gameReminder := game_reminder.New(gameReminderConfig)

		lotteryReminderRabbitMQProducerConfig := rabbitmqproducer.Config{
			QueueName:       viper.GetString("remind_manager.lottery.queue.name"),
			RabbitMQChannel: rabbitMQChannel,
		}
		lotteryReminderRabbitMQProducer := rabbitmqproducer.New(lotteryReminderRabbitMQProducerConfig)

		if err := lotteryReminderRabbitMQProducer.Connect(ctx); err != nil {
			return fmt.Errorf("lottery reminder producer connect error: %w", err)
		}

		lotteryReminderConfig := lottery_reminder.Config{
			Croupier:          croupier,
			GamePlayersFacade: gamePlayersFacade,
			RabbitMQProducer:  lotteryReminderRabbitMQProducer,
		}
		lotteryReminder := lottery_reminder.New(lotteryReminderConfig)

		remindManagerConfig := remindmanager.Config{
			Reminders: []remindmanager.Reminder{
				gameReminder,
				lotteryReminder,
			},
		}
		remindManager := remindmanager.New(remindManagerConfig)

		logger.Infof(ctx, "starting remind manager")
		return remindManager.Start(ctx)
	})

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			if setting.Key == "vcs.revision" {
				logger.InfoKV(ctx, "application started", zap.String("vcs.revision", setting.Value))
			}
		}
	}

	if err := g.Wait(); err != nil {
		logger.Fatal(ctx, err)
	}
}
