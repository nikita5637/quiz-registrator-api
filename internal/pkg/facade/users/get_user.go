package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	telegram_utils "github.com/nikita5637/quiz-registrator-api/utils/telegram"
	"google.golang.org/grpc/metadata"
)

// GetUser ...
func (f *Facade) GetUser(ctx context.Context) (model.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return model.User{}, nil
	}

	telegramIDs := md.Get(telegram_utils.TelegramClientID)
	if len(telegramIDs) == 0 {
		return model.User{}, nil
	}

	telegramIDStr := telegramIDs[0]
	telegramID, err := strconv.ParseInt(telegramIDStr, 10, 64)
	if err != nil {
		return model.User{}, fmt.Errorf("get user error: %w", err)
	}

	user, err := f.userStorage.GetUserByTelegramID(ctx, telegramID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("get user error: %w", model.ErrUserNotFound)
		}

		return model.User{}, fmt.Errorf("get user error: %w", err)
	}

	return user, nil
}

// GetUserByID ...
func (f *Facade) GetUserByID(ctx context.Context, userID int32) (model.User, error) {
	user, err := f.userStorage.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("get user by ID error: %w", model.ErrUserNotFound)
		}

		return model.User{}, fmt.Errorf("get user by ID error: %w", err)
	}

	return user, nil
}

// GetUserByTelegramID ...
func (f *Facade) GetUserByTelegramID(ctx context.Context, telegramID int64) (model.User, error) {
	user, err := f.userStorage.GetUserByTelegramID(ctx, telegramID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf("get user by telegram ID error: %w", model.ErrUserNotFound)
		}

		return model.User{}, fmt.Errorf("get user by telegram ID error: %w", err)
	}

	return user, nil
}
