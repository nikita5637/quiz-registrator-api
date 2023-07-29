package lottery

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/remind-manager/lottery/mocks"
)

type fixture struct {
	ctx      context.Context
	reminder *Reminder

	croupier          *mocks.Croupier
	gamePlayersFacade *mocks.GamePlayersFacade
	rabbitMQProducer  *mocks.RabbitMQProducer
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		croupier:          mocks.NewCroupier(t),
		gamePlayersFacade: mocks.NewGamePlayersFacade(t),
		rabbitMQProducer:  mocks.NewRabbitMQProducer(t),
	}

	fx.reminder = &Reminder{
		alreadyRemindedGames: make(map[int32]struct{}),
		croupier:             fx.croupier,
		gamePlayersFacade:    fx.gamePlayersFacade,
		rabbitMQProducer:     fx.rabbitMQProducer,
	}

	t.Cleanup(func() {})

	return fx
}
