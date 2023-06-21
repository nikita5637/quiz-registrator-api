package admin

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/admin/mocks"
)

type fixture struct {
	ctx context.Context

	userRolesFacade *mocks.UserRolesFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		userRolesFacade: mocks.NewUserRolesFacade(t),
	}

	fx.implementation = New(Config{
		UserRolesFacade: fx.userRolesFacade,
	})

	t.Cleanup(func() {})

	return fx
}
