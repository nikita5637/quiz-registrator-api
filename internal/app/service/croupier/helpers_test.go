package croupier

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/croupier/mocks"
)

type fixture struct {
	ctx            context.Context
	implementation *Implemintation

	croupier *mocks.Croupier

	gamePlayersFacade *mocks.GamePlayersFacade
	gamesFacade       *mocks.GamesFacade
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		croupier: mocks.NewCroupier(t),

		gamePlayersFacade: mocks.NewGamePlayersFacade(t),
		gamesFacade:       mocks.NewGamesFacade(t),
	}

	fx.implementation = &Implemintation{
		croupier: fx.croupier,

		gamePlayersFacade: fx.gamePlayersFacade,
		gamesFacade:       fx.gamesFacade,
	}

	t.Cleanup(func() {})

	return fx
}
