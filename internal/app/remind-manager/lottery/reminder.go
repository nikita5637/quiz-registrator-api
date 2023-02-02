//go:generate mockery --case underscore --name Croupier --with-expecter
//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQChannel --with-expecter

package game

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/reminder"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Croupier ...
type Croupier interface {
	GetGamesWithActiveLottery(ctx context.Context) ([]model.Game, error)
}

// GamesFacade ...
type GamesFacade interface {
	GetPlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
}

// RabbitMQChannel ...
type RabbitMQChannel interface {
	PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
}

// Reminder ...
type Reminder struct {
	alreadyRemindedGames map[int32]struct{}
	croupier             Croupier
	gamesFacade          GamesFacade
	queueName            string
	rabbitMQChannel      RabbitMQChannel
}

// Config ...
type Config struct {
	Croupier        Croupier
	GamesFacade     GamesFacade
	QueueName       string
	RabbitMQChannel RabbitMQChannel
}

// New ...
func New(cfg Config) *Reminder {
	return &Reminder{
		alreadyRemindedGames: make(map[int32]struct{}, 0),
		croupier:             cfg.Croupier,
		gamesFacade:          cfg.GamesFacade,
		queueName:            cfg.QueueName,
		rabbitMQChannel:      cfg.RabbitMQChannel,
	}
}

// Run ...
func (r *Reminder) Run(ctx context.Context) error {
	logger.Info(ctx, "starting lottery reminder")

	err := r.run(ctx)
	if err != nil {
		logger.Errorf(ctx, "lottery reminder error: %s", err.Error())
		return err
	}

	logger.Info(ctx, "lottery reminder done")

	return nil
}

func (r *Reminder) run(ctx context.Context) error {
	_, err := r.rabbitMQChannel.QueueDeclare(
		r.queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue declare error: %w", err)
	}

	games, err := r.croupier.GetGamesWithActiveLottery(ctx)
	if err != nil {
		return fmt.Errorf("get games with active lottery error: %w", err)
	}

	for _, game := range games {
		if _, ok := r.alreadyRemindedGames[game.ID]; ok {
			continue
		}

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

		remind := reminder.Lottery{
			GameID:    game.ID,
			PlayerIDs: playerIDs,
		}

		body, err := json.Marshal(remind)
		if err != nil {
			logger.Errorf(ctx, "remind marshal error: %s", err.Error())
			continue
		}

		err = r.rabbitMQChannel.PublishWithContext(ctx,
			"",
			r.queueName,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			logger.Errorf(ctx, "message publish error: %s", err.Error())
			continue
		}

		r.alreadyRemindedGames[game.ID] = struct{}{}
	}

	return nil
}
