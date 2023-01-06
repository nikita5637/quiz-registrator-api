package games

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
)

type fixture struct {
	ctx    context.Context
	facade *Facade

	gameStorage       *mocks.GameStorage
	gamePlayerStorage *mocks.GamePlayerStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		gameStorage:       mocks.NewGameStorage(t),
		gamePlayerStorage: mocks.NewGamePlayerStorage(t),
	}

	fx.facade = NewFacade(Config{
		GameStorage:       fx.gameStorage,
		GamePlayerStorage: fx.gamePlayerStorage,
	})

	t.Cleanup(func() {})

	return fx
}
