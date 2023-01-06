package places

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
)

type fixture struct {
	ctx    context.Context
	facade *Facade

	placeStorage *mocks.PlaceStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		placeStorage: mocks.NewPlaceStorage(t),
	}

	fx.facade = &Facade{
		placeStorage: fx.placeStorage,
	}

	t.Cleanup(func() {
	})

	return fx
}
