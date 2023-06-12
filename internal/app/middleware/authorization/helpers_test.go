package authorization

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/middleware/authorization/mocks"
)

type fixture struct {
	ctx context.Context

	userRolesFacade *mocks.UserRolesFacade

	middleware *Middleware
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		userRolesFacade: mocks.NewUserRolesFacade(t),
	}

	fx.middleware = New(Config{
		UserRolesFacade: fx.userRolesFacade,
	})

	t.Cleanup(func() {})

	return fx
}
