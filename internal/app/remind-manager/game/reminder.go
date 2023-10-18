//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name GamePlayersFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQProducer --with-expecter

package game

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/reminder"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GamesFacade ...
type GamesFacade interface {
	GetTodaysGames(ctx context.Context) ([]model.Game, error)
}

// GamePlayersFacade ...
type GamePlayersFacade interface { // nolint:revive
	GetGamePlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
}

// RabbitMQProducer ...
type RabbitMQProducer interface {
	Send(ctx context.Context, message interface{}) error
}

// Reminder ...
type Reminder struct {
	gamesFacade       GamesFacade
	gamePlayersFacade GamePlayersFacade
	rabbitMQProducer  RabbitMQProducer
}

// Config ...
type Config struct {
	GamesFacade       GamesFacade
	GamePlayersFacade GamePlayersFacade
	RabbitMQProducer  RabbitMQProducer
}

// New ...
func New(cfg Config) *Reminder {
	return &Reminder{
		gamesFacade:       cfg.GamesFacade,
		gamePlayersFacade: cfg.GamePlayersFacade,
		rabbitMQProducer:  cfg.RabbitMQProducer,
	}
}

// Run runs at 07:00 UTC
func (r *Reminder) Run(ctx context.Context) error {
	if timeutils.TimeNow().UTC().Hour() == viper.GetInt("remind_manager.game.hour") &&
		timeutils.TimeNow().UTC().Minute() == viper.GetInt("remind_manager.game.minute") {
		ctx = logger.ToContext(ctx, logger.FromContext(ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "game reminder"),
		)))

		logger.Info(ctx, "starting reminder")

		err := r.run(ctx)
		if err != nil {
			logger.Errorf(ctx, "reminder error: %s", err.Error())
			return err
		}

		logger.Info(ctx, "reminder done")
	}

	return nil
}

func (r *Reminder) run(ctx context.Context) error {
	games, err := r.gamesFacade.GetTodaysGames(ctx)
	if err != nil {
		return fmt.Errorf("get todays games error: %w", err)
	}

	if len(games) == 0 {
		logger.Info(ctx, "there are not todays games")
		return nil
	}

	for _, game := range games {
		players, err := r.gamePlayersFacade.GetGamePlayersByGameID(ctx, game.ID)
		if err != nil {
			logger.ErrorKV(ctx, "get players by game ID error", "error", err, "gameID", game.ID)
			continue
		}

		playerIDs := make([]int32, 0, len(players))
		for _, player := range players {
			if player.UserID.IsPresent() {
				playerIDs = append(playerIDs, player.UserID.Value())
			}
		}

		if len(playerIDs) == 0 {
			logger.WarnKV(ctx, "there are not players to remind", "gameID", game.ID)
			continue
		}

		reminder := reminder.Game{
			GameID:    game.ID,
			PlayerIDs: playerIDs,
		}

		err = r.rabbitMQProducer.Send(ctx, reminder)
		if err != nil {
			logger.Errorf(ctx, "send message error: %s", err.Error())
			continue
		}

		logger.InfoKV(ctx, "message published", "reminder", reminder)
	}

	return nil
}
