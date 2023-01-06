package mysql

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// UserStorageAdapter ...
type UserStorageAdapter struct {
	userStorage *UserStorage
}

// NewUserStorageAdapter ...
func NewUserStorageAdapter(db *sql.DB) *UserStorageAdapter {
	return &UserStorageAdapter{
		userStorage: NewUserStorage(db),
	}
}

// GetUserByID ...
func (a *UserStorageAdapter) GetUserByID(ctx context.Context, userID int32) (model.User, error) {
	user, err := a.userStorage.GetUserByID(ctx, int(userID))
	if err != nil {
		return model.User{}, err
	}

	return convertDBUserToModelUser(*user), nil
}

// GetUserByTelegramID ...
func (a *UserStorageAdapter) GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error) {
	user, err := a.userStorage.GetUserByTelegramID(ctx, telegramID)
	if err != nil {
		return model.User{}, err
	}

	return convertDBUserToModelUser(*user), nil
}

// Insert ...
func (a *UserStorageAdapter) Insert(ctx context.Context, user model.User) (int32, error) {
	id, err := a.userStorage.Insert(ctx, convertModelUserToDBUser(user))
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// Update ...
func (a *UserStorageAdapter) Update(ctx context.Context, user model.User) error {
	return a.userStorage.Update(ctx, convertModelUserToDBUser(user))
}

func convertDBUserToModelUser(user User) model.User {
	return model.User{
		ID:         int32(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		Email:      user.Email.String,
		Phone:      user.Phone.String,
		State:      model.UserState(user.State),
		CreatedAt:  model.DateTime(user.CreatedAt.Time),
		UpdatedAt:  model.DateTime(user.UpdatedAt.Time),
	}
}

func convertModelUserToDBUser(user model.User) User {
	ret := User{
		ID:         int(user.ID),
		Name:       user.Name,
		TelegramID: user.TelegramID,
		State:      int(user.State),
	}

	if len(user.Email) > 0 {
		ret.Email = sql.NullString{
			String: user.Email,
			Valid:  true,
		}
	}

	if len(user.Phone) > 0 {
		ret.Phone = sql.NullString{
			String: user.Phone,
			Valid:  true,
		}
	}

	return ret
}
