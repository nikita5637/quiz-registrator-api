//go:generate mockery --case underscore --name UsersFacade --with-expecter

package usermanager

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
)

// UsersFacade ...
type UsersFacade interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUser(ctx context.Context, userID int32) (model.User, error)
	GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error)
	PatchUser(ctx context.Context, user model.User) (model.User, error)
}

// Implementation ...
type Implementation struct {
	usersFacade UsersFacade

	usermanagerpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	UsersFacade UsersFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		usersFacade: cfg.UsersFacade,
	}
}
