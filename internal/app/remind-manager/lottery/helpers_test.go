package game

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/lottery/mocks"
)

type fixture struct {
	ctx      context.Context
	reminder *Reminder

	croupier        *mocks.Croupier
	gamesFacade     *mocks.GamesFacade
	rabbitMQChannel *mocks.RabbitMQChannel
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		croupier:        mocks.NewCroupier(t),
		gamesFacade:     mocks.NewGamesFacade(t),
		rabbitMQChannel: mocks.NewRabbitMQChannel(t),
	}

	fx.reminder = &Reminder{
		alreadyRemindedGames: make(map[int32]struct{}),
		croupier:             fx.croupier,
		gamesFacade:          fx.gamesFacade,
		queueName:            "queue",
		rabbitMQChannel:      fx.rabbitMQChannel,
	}

	t.Cleanup(func() {})

	return fx
}