package gamephotos

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
)

type fixture struct {
	ctx    context.Context
	facade *Facade

	gameStorage       *mocks.GameStorage
	gamePhotoStorage  *mocks.GamePhotoStorage
	gameResultStorage *mocks.GameResultStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gameStorage:       mocks.NewGameStorage(t),
		gamePhotoStorage:  mocks.NewGamePhotoStorage(t),
		gameResultStorage: mocks.NewGameResultStorage(t),
	}

	fx.facade = &Facade{
		gameStorage:       fx.gameStorage,
		gamePhotoStorage:  fx.gamePhotoStorage,
		gameResultStorage: fx.gameResultStorage,
	}

	t.Cleanup(func() {})

	return fx
}
