package place

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/place/mocks"
)

type fixture struct {
	ctx            context.Context
	implementation *Implementation

	placesFacade *mocks.PlacesFacade
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		placesFacade: mocks.NewPlacesFacade(t),
	}

	fx.implementation = &Implementation{
		placesFacade: fx.placesFacade,
	}

	t.Cleanup(func() {})

	return fx
}
