package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetUser ...
func (f *Facade) GetUser(ctx context.Context, userID int32) (model.User, error) {
	modelUser := model.User{}
	err := f.db.RunTX(ctx, "GetUser", func(ctx context.Context) error {
		dbUser, err := f.userStorage.GetUserByID(ctx, int(userID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get user by ID error: %w", ErrUserNotFound)
			}

			return fmt.Errorf("get user by ID error: %w", err)
		}

		modelUser = convertDBUserToModelUser(*dbUser)

		return nil
	})
	if err != nil {
		return model.User{}, err
	}

	return modelUser, nil
}

// GetUserByTelegramID ...
func (f *Facade) GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error) {
	modelUser := model.User{}
	err := f.db.RunTX(ctx, "GetUserByTelegramID", func(ctx context.Context) error {
		dbUser, err := f.userStorage.GetUserByTelegramID(ctx, telegramID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get user by Telegram ID error: %w", ErrUserNotFound)
			}

			return fmt.Errorf("get user by Telegram ID error: %w", err)
		}

		modelUser = convertDBUserToModelUser(*dbUser)

		return nil
	})
	if err != nil {
		return model.User{}, err
	}

	return modelUser, nil
}
