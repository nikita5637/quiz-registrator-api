package photomanager

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/photo_manager/mocks"
)

type fixture struct {
	ctx context.Context

	gamePhotosFacade *mocks.GamePhotosFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gamePhotosFacade: mocks.NewGamePhotosFacade(t),
	}

	fx.implementation = &Implementation{
		gamePhotosFacade: fx.gamePhotosFacade,
	}

	t.Cleanup(func() {})

	return fx
}
