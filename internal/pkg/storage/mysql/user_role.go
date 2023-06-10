package mysql

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// UserRoleStorageAdapter ...
type UserRoleStorageAdapter struct {
	userRoleStorage *UserRoleStorage
}

// NewUserRoleStorageAdapter ...
func NewUserRoleStorageAdapter(txManager *tx.Manager) *UserRoleStorageAdapter {
	return &UserRoleStorageAdapter{
		userRoleStorage: NewUserRoleStorage(txManager),
	}
}

// DeleteUserRole ...
func (a *UserRoleStorageAdapter) DeleteUserRole(ctx context.Context, userRoleID int) error {
	return a.userRoleStorage.Delete(ctx, userRoleID)
}

// GetUserRoleByID ...
func (a *UserRoleStorageAdapter) GetUserRoleByID(ctx context.Context, userRoleID int) (*UserRole, error) {
	return a.userRoleStorage.GetUserRoleByID(ctx, userRoleID)
}

// GetUserRoles ...
func (a *UserRoleStorageAdapter) GetUserRoles(ctx context.Context) ([]UserRole, error) {
	return a.userRoleStorage.Find(ctx, builder.IsNull{
		"deleted_at",
	}, "id")
}

// GetUserRolesByUserID ...
func (a *UserRoleStorageAdapter) GetUserRolesByUserID(ctx context.Context, userID int) ([]UserRole, error) {
	return a.userRoleStorage.Find(ctx, builder.And(
		builder.Eq{
			"fk_user_id": userID,
		},
		builder.IsNull{
			"deleted_at",
		}),
		"id")
}

// Insert ...
func (a *UserRoleStorageAdapter) Insert(ctx context.Context, userRole UserRole) (int, error) {
	id, err := a.userRoleStorage.Insert(ctx, userRole)
	if err != nil {
		return 0, err
	}

	return id, nil
}
