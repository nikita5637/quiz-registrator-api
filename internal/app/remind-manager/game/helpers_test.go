package game

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/game/mocks"
)

type fixture struct {
	ctx      context.Context
	reminder *Reminder

	gamesFacade      *mocks.GamesFacade
	rabbitMQProducer *mocks.RabbitMQProducer
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade:      mocks.NewGamesFacade(t),
		rabbitMQProducer: mocks.NewRabbitMQProducer(t),
	}

	fx.reminder = &Reminder{
		gamesFacade:      fx.gamesFacade,
		rabbitMQProducer: fx.rabbitMQProducer,
	}

	t.Cleanup(func() {})

	return fx
}
