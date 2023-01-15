package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/app/registrator"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/croupier/quiz_please"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/elasticsearch"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gamephotos"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"go.uber.org/zap"
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
		panic(err)
	}
	defer db.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		oscall := <-c
		logger.Infof(ctx, "system call recieved: %+v", oscall)
		cancel()
	}()

	croupierConfig := croupier.Config{
		QuizPleaseCroupier: quiz_please.New(),
	}

	croupier := croupier.New(croupierConfig)

	txManager := tx.NewManager(db)

	gameStorage := storage.NewGameStorage(txManager)
	gamePhotoStorage := storage.NewGamePhotoStorage(txManager)
	gamePlayerStorage := storage.NewGamePlayerStorage(txManager)
	gameResultStorage := storage.NewGameResultStorage(txManager)
	leagueStorage := storage.NewLeagueStorage(txManager)
	placeStorage := storage.NewPlaceStorage(txManager)
	userStorage := storage.NewUserStorage(txManager)

	gamesFacadeConfig := games.Config{
		GamePlayerStorage: gamePlayerStorage,
		GameStorage:       gameStorage,
		TxManager:         txManager,
	}

	gamesFacade := games.NewFacade(gamesFacadeConfig)

	gamePhotosFacadeConfig := gamephotos.Config{
		GameStorage:       gameStorage,
		GamePhotoStorage:  gamePhotoStorage,
		GameResultStorage: gameResultStorage,
		TxManager:         txManager,
	}
	gamePhotosFacade := gamephotos.NewFacade(gamePhotosFacadeConfig)

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

		GamesFacade:      gamesFacade,
		GamePhotosFacade: gamePhotosFacade,
		LeaguesFacade:    leaguesFacade,
		PlacesFacade:     placesFacade,
		UsersFacade:      usersFacade,
	}

	reg := registrator.New(registratorConfig)

	err = reg.ListenAndServe(ctx)
	if err != nil {
		logger.Panic(ctx, err)
	}
}
