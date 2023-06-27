package registrator

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/registrator/mocks"
)

type fixture struct {
	ctx         context.Context
	registrator *Registrator

	gamesFacade  *mocks.GamesFacade
	placesFacade *mocks.PlacesFacade
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamesFacade:  mocks.NewGamesFacade(t),
		placesFacade: mocks.NewPlacesFacade(t),
	}

	fx.registrator = &Registrator{
		gamesFacade:  fx.gamesFacade,
		placesFacade: fx.placesFacade,
	}

	t.Cleanup(func() {})

	return fx
}
