//go:generate mockery --case underscore --name UsersFacade --with-expecter

package authentication

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

const (
	authenticationTypeTelegramID  = "authentication_type_telegram_id"
	authenticationTypeServiceName = "authentication_type_service_name"

	moduleNameHeader       = "x-module-name"
	serviceNameHeader      = "x-service-name"
	telegramClientIDHeader = "x-telegram-client-id"
)

// UsersFacade ...
type UsersFacade interface {
	GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error)
}

// Middleware ...
type Middleware struct {
	usersFacade UsersFacade
}

// Config ...
type Config struct {
	UsersFacade UsersFacade
}

// New ...
func New(cfg Config) *Middleware {
	return &Middleware{
		usersFacade: cfg.UsersFacade,
	}
}
