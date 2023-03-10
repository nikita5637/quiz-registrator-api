//go:generate mockery --case underscore --name UserStorage --with-expecter

package storage

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// UserStorage ...
type UserStorage interface {
	GetUserByID(ctx context.Context, userID int32) (model.User, error)
	GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error)
	Insert(ctx context.Context, user model.User) (int32, error)
	Update(ctx context.Context, user model.User) error
}

// NewUserStorage ...
func NewUserStorage(db *sql.DB) UserStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewUserStorageAdapter(db)
	}

	return nil
}
