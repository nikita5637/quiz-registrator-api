//go:generate mockery --case underscore --name UserRolesFacade --with-expecter

package admin

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
)

// UserRolesFacade ...
type UserRolesFacade interface {
	CreateUserRole(ctx context.Context, userRole model.UserRole) (model.UserRole, error)
	DeleteUserRole(ctx context.Context, id int32) error
	GetUserRolesByUserID(ctx context.Context, userID int32) ([]model.UserRole, error)
	ListUserRoles(ctx context.Context) ([]model.UserRole, error)
}

// Implementation ...
type Implementation struct {
	userRolesFacade UserRolesFacade

	adminpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	UserRolesFacade UserRolesFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		userRolesFacade: cfg.UserRolesFacade,
	}
}
