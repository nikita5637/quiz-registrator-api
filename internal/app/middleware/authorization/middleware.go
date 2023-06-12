//go:generate mockery --case underscore --name UserRolesFacade --with-expecter

package authorization

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// UserRolesFacade ...
type UserRolesFacade interface {
	GetUserRolesByUserID(ctx context.Context, userID int32) ([]model.UserRole, error)
}

// Middleware ...
type Middleware struct {
	userRolesFacade UserRolesFacade
}

// Config ...
type Config struct {
	UserRolesFacade UserRolesFacade
}

// New ...
func New(cfg Config) *Middleware {
	return &Middleware{
		userRolesFacade: cfg.UserRolesFacade,
	}
}
