//go:generate mockery --case underscore --name Croupier --with-expecter
//go:generate mockery --case underscore --name GamePlayersFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQProducer --with-expecter

package lottery

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/reminder"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"go.uber.org/zap"
)

// Croupier ...
type Croupier interface {
	GetGamesWithActiveLottery(ctx context.Context) ([]model.Game, error)
}

// GamePlayersFacade ...
type GamePlayersFacade interface {
	GetGamePlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
}

// RabbitMQProducer ...
type RabbitMQProducer interface {
	Send(ctx context.Context, message interface{}) error
}

// Reminder ...
type Reminder struct {
	alreadyRemindedGames map[int32]struct{}
	croupier             Croupier
	gamePlayersFacade    GamePlayersFacade
	rabbitMQProducer     RabbitMQProducer
}

// Config ...
type Config struct {
	Croupier          Croupier
	GamePlayersFacade GamePlayersFacade
	RabbitMQProducer  RabbitMQProducer
}

// New ...
func New(cfg Config) *Reminder {
	return &Reminder{
		alreadyRemindedGames: make(map[int32]struct{}, 0),
		croupier:             cfg.Croupier,
		gamePlayersFacade:    cfg.GamePlayersFacade,
		rabbitMQProducer:     cfg.RabbitMQProducer,
	}
}

// Run ...
func (r *Reminder) Run(ctx context.Context) error {
	if timeutils.TimeNow().UTC().Minute() == 0 ||
		timeutils.TimeNow().UTC().Minute() == 30 {
		ctx = logger.ToContext(ctx, logger.FromContext(ctx).WithOptions(zap.Fields(
			zap.String("reminder_name", "lottery reminder"),
		)))

		logger.Info(ctx, "starting reminder")

		err := r.run(ctx)
		if err != nil {
			logger.ErrorKV(ctx, "running reminder error", zap.Error(err))
			return err
		}

		logger.Info(ctx, "reminder done")
	}

	return nil
}

func (r *Reminder) run(ctx context.Context) error {
	games, err := r.croupier.GetGamesWithActiveLottery(ctx)
	if err != nil {
		return fmt.Errorf("get games with active lottery error: %w", err)
	}

	if len(games) == 0 {
		logger.Info(ctx, "there are not games with active lottery")
		return nil
	}

	for _, game := range games {
		if _, ok := r.alreadyRemindedGames[game.ID]; ok {
			continue
		}

		players, err := r.gamePlayersFacade.GetGamePlayersByGameID(ctx, game.ID)
		if err != nil {
			logger.ErrorKV(ctx, "getting game players by game ID error", zap.Error(err), "gameID", game.ID)
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

		reminder := reminder.Lottery{
			GameID:    game.ID,
			LeagueID:  game.LeagueID,
			PlayerIDs: playerIDs,
		}

		err = r.rabbitMQProducer.Send(ctx, reminder)
		if err != nil {
			logger.ErrorKV(ctx, "sending message error", zap.Error(err))
			continue
		}

		logger.InfoKV(ctx, "message published", "reminder", reminder)

		r.alreadyRemindedGames[game.ID] = struct{}{}
	}

	return nil
}
