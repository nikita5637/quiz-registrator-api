package game

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/game/mocks"
)

type fixture struct {
	ctx      context.Context
	reminder *Reminder

	gamesFacade     *mocks.GamesFacade
	rabbitMQChannel *mocks.RabbitMQChannel
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade:     mocks.NewGamesFacade(t),
		rabbitMQChannel: mocks.NewRabbitMQChannel(t),
	}

	fx.reminder = &Reminder{
		gamesFacade:     fx.gamesFacade,
		queueName:       "queue",
		rabbitMQChannel: fx.rabbitMQChannel,
	}

	t.Cleanup(func() {})

	return fx
}
