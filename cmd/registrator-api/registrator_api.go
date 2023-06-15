package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	adminservice "github.com/nikita5637/quiz-registrator-api/internal/app/admin_service"
	certificatemanager "github.com/nikita5637/quiz-registrator-api/internal/app/certificate_manager"
	gameresultmanager "github.com/nikita5637/quiz-registrator-api/internal/app/game_result_manager"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authentication"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authorization"
	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/log"
	"github.com/nikita5637/quiz-registrator-api/internal/app/registrator"
	remindmanager "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager"
	game_reminder "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/game"
	lottery_reminder "github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/lottery"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/quiz_please"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/squiz"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/elasticsearch"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gamephotos"
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
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./config.toml", "path to config file")
}

func main() {
	flag.Parse()

	ctx := context.Background()

	var err error
	err = config.ParseConfigFile(configPath)
	if err != nil {
		panic(err)
	}

	logsCombiner := &logger.Combiner{}
	logsCombiner = logsCombiner.WithWriter(os.Stdout)

	elasticLogsEnabled := config.GetValue("ElasticLogsEnabled").Bool()
	if elasticLogsEnabled {
		var elasticClient *elasticsearch.Client
		elasticClient, err = elasticsearch.New(elasticsearch.Config{
			ElasticAddress: config.GetElasticAddress(),
			ElasticIndex:   config.GetValue("ElasticIndex").String(),
		})
		if err != nil {
			panic(err)
		}

		logger.Info(ctx, "initialized elasticsearch client")
		logsCombiner = logsCombiner.WithWriter(elasticClient)
	}

	logLevel := config.GetLogLevel()
	logger.SetGlobalLogger(logger.NewLogger(logLevel, logsCombiner, zap.Fields(
		zap.String("module", "registrator-api"),
	)))
	logger.InfoKV(ctx, "initialized logger", "log level", logLevel)

	db, err := storage.NewDB()
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		oscall := <-c
		logger.Infof(ctx, "system call recieved: %+v", oscall)
		cancel()
	}()

	txManager := tx.NewManager(db)

	gameStorage := storage.NewGameStorage(txManager)
	gamePlayerStorage := storage.NewGamePlayerStorage(txManager)

	icsRabbitMQProducerConfig := rabbitmqproducer.Config{
		QueueName:       config.GetValue("RabbitMQICSQueueName").String(),
		RabbitMQChannel: rabbitMQChannel,
	}
	icsRabbitMQProducer := rabbitmqproducer.New(icsRabbitMQProducerConfig)

	if err := icsRabbitMQProducer.Connect(ctx); err != nil {
		logger.Fatalf(ctx, "ICS producer connect error: %s", err.Error())
	}

	gamesFacadeConfig := games.Config{
		GameStorage:       gameStorage,
		GamePlayerStorage: gamePlayerStorage,
		RabbitMQProducer:  icsRabbitMQProducer,
		TxManager:         txManager,
	}

	gamesFacade := games.NewFacade(gamesFacadeConfig)

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
		certificateStorage := storage.NewCertificateStorage(txManager)
		gamePhotoStorage := storage.NewGamePhotoStorage(txManager)
		gameResultStorage := storage.NewGameResultStorage(txManager)
		leagueStorage := storage.NewLeagueStorage(txManager)
		placeStorage := storage.NewPlaceStorage(txManager)
		userStorage := storage.NewUserStorage(txManager)
		userRoleStorage := storage.NewUserRoleStorage(txManager)

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

		certificateManagerConfig := certificatemanager.Config{
			CertificatesFacade: certificatesFacade,
		}
		certificateManager := certificatemanager.New(certificateManagerConfig)

		gamePhotosFacadeConfig := gamephotos.Config{
			GameStorage:       gameStorage,
			GamePhotoStorage:  gamePhotoStorage,
			GameResultStorage: gameResultStorage,
			TxManager:         txManager,
		}
		gamePhotosFacade := gamephotos.NewFacade(gamePhotosFacadeConfig)

		gameResultsFacadeConfig := gameresults.Config{
			GameResultStorage: gameResultStorage,
			TxManager:         txManager,
		}
		gameResultsFacade := gameresults.NewFacade(gameResultsFacadeConfig)

		gameResultManagerConfig := gameresultmanager.Config{
			GameResultsFacade: gameResultsFacade,
		}
		gameResultManager := gameresultmanager.New(gameResultManagerConfig)

		leaguesFacadeConfig := leagues.Config{
			LeagueStorage: leagueStorage,
		}
		leaguesFacade := leagues.NewFacade(leaguesFacadeConfig)

		placesFacadeConfig := places.Config{
			PlaceStorage: placeStorage,
		}
		placesFacade := places.NewFacade(placesFacadeConfig)

		usersFacadeConfig := users.Config{
			UserStorage: userStorage,
		}
		usersFacade := users.NewFacade(usersFacadeConfig)

		authenticationMiddleware := authentication.New(authentication.Config{
			UsersFacade: usersFacade,
		})

		authorizationMiddleware := authorization.New(authorization.Config{
			UserRolesFacade: userRolesFacade,
		})

		logMiddleware := log.New()

		registratorConfig := registrator.Config{
			BindAddr: config.GetBindAddress(),

			Croupier: croupier,

			AuthenticationMiddleware: authenticationMiddleware,
			AuthorizationMiddleware:  authorizationMiddleware,
			LogMiddleware:            logMiddleware,

			AdminService:       adminService,
			CertificateManager: certificateManager,
			GameResultManager:  gameResultManager,

			GamesFacade:      gamesFacade,
			GamePhotosFacade: gamePhotosFacade,
			LeaguesFacade:    leaguesFacade,
			PlacesFacade:     placesFacade,
			UsersFacade:      usersFacade,
		}

		reg := registrator.New(registratorConfig)

		logger.Infof(ctx, "starting registrator")
		return reg.ListenAndServe(ctx)
	})

	g.Go(func() error {
		gameReminderRabbitMQProducerConfig := rabbitmqproducer.Config{
			QueueName:       config.GetValue("RabbitMQGameReminderQueueName").String(),
			RabbitMQChannel: rabbitMQChannel,
		}
		gameReminderRabbitMQProducer := rabbitmqproducer.New(gameReminderRabbitMQProducerConfig)

		if err := gameReminderRabbitMQProducer.Connect(ctx); err != nil {
			logger.Fatalf(ctx, "game reminder producer connect error: %s", err.Error())
		}

		gameReminderConfig := game_reminder.Config{
			GamesFacade:      gamesFacade,
			RabbitMQProducer: gameReminderRabbitMQProducer,
		}
		gameReminder := game_reminder.New(gameReminderConfig)

		lotteryReminderRabbitMQProducerConfig := rabbitmqproducer.Config{
			QueueName:       config.GetValue("RabbitMQLotteryReminderQueueName").String(),
			RabbitMQChannel: rabbitMQChannel,
		}
		lotteryReminderRabbitMQProducer := rabbitmqproducer.New(lotteryReminderRabbitMQProducerConfig)

		if err := lotteryReminderRabbitMQProducer.Connect(ctx); err != nil {
			logger.Fatalf(ctx, "lottery reminder producer connect error: %s", err.Error())
		}

		lotteryReminderConfig := lottery_reminder.Config{
			Croupier:         croupier,
			GamesFacade:      gamesFacade,
			RabbitMQProducer: lotteryReminderRabbitMQProducer,
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

	if err := g.Wait(); err != nil {
		logger.Fatal(ctx, err)
	}
}
