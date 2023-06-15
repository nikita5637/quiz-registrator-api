//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQProducer --with-expecter

package game

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/reminder"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"go.uber.org/zap"
)

// GamesFacade ...
type GamesFacade interface {
	GetPlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
	GetTodaysGames(ctx context.Context) ([]model.Game, error)
}

// RabbitMQProducer ...
type RabbitMQProducer interface {
	Send(ctx context.Context, message interface{}) error
}

// Reminder ...
type Reminder struct {
	gamesFacade      GamesFacade
	rabbitMQProducer RabbitMQProducer
}

// Config ...
type Config struct {
	GamesFacade      GamesFacade
	RabbitMQProducer RabbitMQProducer
}

// New ...
func New(cfg Config) *Reminder {
	return &Reminder{
		gamesFacade:      cfg.GamesFacade,
		rabbitMQProducer: cfg.RabbitMQProducer,
	}
}

// Run runs at 10:00
func (r *Reminder) Run(ctx context.Context) error {
	if time_utils.TimeNow().Hour() == 10 &&
		time_utils.TimeNow().Minute() == 0 {
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
		players, err := r.gamesFacade.GetPlayersByGameID(ctx, game.ID)
		if err != nil {
			logger.ErrorKV(ctx, "get players by game ID error", "error", err, "gameID", game.ID)
			continue
		}

		playerIDs := make([]int32, 0, len(players))
		for _, player := range players {
			if player.FkUserID != 0 {
				playerIDs = append(playerIDs, player.FkUserID)
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
