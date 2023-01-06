package users

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
)

type fixture struct {
	ctx    context.Context
	facade *Facade

	userStorage *mocks.UserStorage
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		userStorage: mocks.NewUserStorage(t),
	}

	fx.facade = NewFacade(Config{
		UserStorage: fx.userStorage,
	})

	t.Cleanup(func() {})

	return fx
}
