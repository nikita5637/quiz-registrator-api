package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	certificatemanager "github.com/nikita5637/quiz-registrator-api/internal/app/certificate_manager"
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
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
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
		logger.Panic(ctx, err)
	}
	defer db.Close()

	rabbitMQConn, err := amqp.Dial(config.GetRabbitMQURL())
	if err != nil {
		logger.Panic(ctx, err)
	}
	defer rabbitMQConn.Close()

	rabbitMQChannel, err := rabbitMQConn.Channel()
	if err != nil {
		logger.Panic(ctx, err)
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

	gamesFacadeConfig := games.Config{
		GameStorage:       gameStorage,
		GamePlayerStorage: gamePlayerStorage,
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

		certificatesFacadeConfig := certificates.Config{
			CertificateStorage: certificateStorage,
			TxManager:          txManager,
		}
		certificatesFacade := certificates.NewFacade(certificatesFacadeConfig)

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

		registratorConfig := registrator.Config{
			BindAddr: config.GetBindAddress(),

			Croupier: croupier,

			CertificateManager: certificateManager,
			GamesFacade:        gamesFacade,
			GamePhotosFacade:   gamePhotosFacade,
			GameResultsFacade:  gameResultsFacade,
			LeaguesFacade:      leaguesFacade,
			PlacesFacade:       placesFacade,
			UsersFacade:        usersFacade,
		}

		reg := registrator.New(registratorConfig)

		logger.Infof(ctx, "starting registrator")
		return reg.ListenAndServe(ctx)
	})

	g.Go(func() error {
		gameReminderQueueName := config.GetValue("RabbitMQGameReminderQueueName").String()
		if gameReminderQueueName == "" {
			return errors.New("empty rabbit MQ game reminder queue name")
		}

		gameReminderConfig := game_reminder.Config{
			GamesFacade:     gamesFacade,
			QueueName:       gameReminderQueueName,
			RabbitMQChannel: rabbitMQChannel,
		}
		gameReminder := game_reminder.New(gameReminderConfig)

		lotteryReminderQueueName := config.GetValue("RabbitMQLotteryReminderQueueName").String()
		if lotteryReminderQueueName == "" {
			return errors.New("empty rabbit MQ lottery reminder queue name")
		}

		lotteryReminderConfig := lottery_reminder.Config{
			Croupier:        croupier,
			GamesFacade:     gamesFacade,
			QueueName:       lotteryReminderQueueName,
			RabbitMQChannel: rabbitMQChannel,
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
		logger.Panic(ctx, err)
	}
}
