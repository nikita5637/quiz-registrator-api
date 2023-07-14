//go:generate mockery --case underscore --name UserStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// UserStorage ...
type UserStorage interface {
	GetUserByID(ctx context.Context, userID int) (*database.User, error)
	GetUserByTelegramID(ctx context.Context, telegramID int64) (*database.User, error)
	Insert(ctx context.Context, user database.User) (int, error)
	Update(ctx context.Context, user database.User) error
}

// NewUserStorage ...
func NewUserStorage(driver string, txManager *tx.Manager) UserStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewUserStorageAdapter(txManager)
	}

	return nil
}
