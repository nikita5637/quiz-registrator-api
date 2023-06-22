//go:generate mockery --case underscore --name UserRoleStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// UserRoleStorage ...
type UserRoleStorage interface {
	DeleteUserRole(ctx context.Context, id int) error
	GetUserRoleByID(ctx context.Context, userRoleID int) (*database.UserRole, error)
	GetUserRoles(ctx context.Context) ([]database.UserRole, error)
	GetUserRolesByUserID(ctx context.Context, userID int) ([]database.UserRole, error)
	Insert(ctx context.Context, userRole database.UserRole) (int, error)
}

// NewUserRoleStorage ...
func NewUserRoleStorage(driver string, txManager *tx.Manager) UserRoleStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewUserRoleStorageAdapter(txManager)
	}

	return nil
}
