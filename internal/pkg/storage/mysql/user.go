package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// UserStorageAdapter ...
type UserStorageAdapter struct {
	userStorage *UserStorage
}

// NewUserStorageAdapter ...
func NewUserStorageAdapter(txManager *tx.Manager) *UserStorageAdapter {
	return &UserStorageAdapter{
		userStorage: NewUserStorage(txManager),
	}
}

// GetUserByID ...
func (a *UserStorageAdapter) GetUserByID(ctx context.Context, userID int) (*User, error) {
	return a.userStorage.GetUserByID(ctx, userID)
}

// GetUserByTelegramID ...
func (a *UserStorageAdapter) GetUserByTelegramID(ctx context.Context, telegramID int64) (*User, error) {
	return a.userStorage.GetUserByTelegramID(ctx, telegramID)
}

// Insert ...
func (a *UserStorageAdapter) Insert(ctx context.Context, user User) (int, error) {
	return a.userStorage.Insert(ctx, user)
}

// Update ...
func (a *UserStorageAdapter) Update(ctx context.Context, user User) error {
	return a.userStorage.Update(ctx, user)
}
