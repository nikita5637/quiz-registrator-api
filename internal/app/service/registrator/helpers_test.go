package registrator

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/registrator/mocks"
)

type fixture struct {
	ctx            context.Context
	implementation *Implementation

	gamesFacade *mocks.GamesFacade
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade: mocks.NewGamesFacade(t),
	}

	fx.implementation = &Implementation{
		gamesFacade: fx.gamesFacade,
	}

	t.Cleanup(func() {})

	return fx
}
