package authentication

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authentication/mocks"
)

type fixture struct {
	ctx context.Context

	usersFacade *mocks.UsersFacade

	middleware *Middleware
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		usersFacade: mocks.NewUsersFacade(t),
	}

	fx.middleware = New(Config{
		UsersFacade: fx.usersFacade,
	})

	t.Cleanup(func() {})

	return fx
}
