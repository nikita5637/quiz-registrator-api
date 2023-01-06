package users

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	telegram_utils "github.com/nikita5637/quiz-registrator-api/utils/telegram"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestFacade_GetUser(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("empty metadata", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.GetUser(fx.ctx)
		assert.Equal(t, model.User{}, got)
		assert.NoError(t, err)
	})

	t.Run("empty telegram client id header value", func(t *testing.T) {
		fx := tearUp(t)

		md := metadata.New(nil)
		ctx := metadata.NewIncomingContext(fx.ctx, md)

		got, err := fx.facade.GetUser(ctx)
		assert.Equal(t, model.User{}, got)
		assert.NoError(t, err)
	})

	t.Run("invalid telegram client id value", func(t *testing.T) {
		fx := tearUp(t)

		md := metadata.New(map[string]string{
			telegram_utils.TelegramClientID: "invalid value",
		})
		ctx := metadata.NewIncomingContext(fx.ctx, md)

		got, err := fx.facade.GetUser(ctx)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
	})

	t.Run("sql.ErrNoRows error", func(t *testing.T) {
		fx := tearUp(t)

		md := metadata.New(map[string]string{
			telegram_utils.TelegramClientID: "1",
		})
		ctx := metadata.NewIncomingContext(fx.ctx, md)

		fx.userStorage.EXPECT().GetUserByTelegramID(ctx, int64(1)).Return(model.User{}, sql.ErrNoRows)

		got, err := fx.facade.GetUser(ctx)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		md := metadata.New(map[string]string{
			telegram_utils.TelegramClientID: "1",
		})
		ctx := metadata.NewIncomingContext(fx.ctx, md)

		fx.userStorage.EXPECT().GetUserByTelegramID(ctx, int64(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.facade.GetUser(ctx)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		md := metadata.New(map[string]string{
			telegram_utils.TelegramClientID: "1",
		})
		ctx := metadata.NewIncomingContext(fx.ctx, md)

		fx.userStorage.EXPECT().GetUserByTelegramID(ctx, int64(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      "email",
			Phone:      "phone",
			State:      1,
			CreatedAt:  model.DateTime(timeNow),
			UpdatedAt:  model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		got, err := fx.facade.GetUser(ctx)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      "email",
			Phone:      "phone",
			State:      1,
			CreatedAt:  model.DateTime(timeNow),
			UpdatedAt:  model.DateTime(timeNow.Add(1 * time.Second)),
		}, got)
		assert.NoError(t, err)
	})
}

func TestFacade_GetUserByTelegramID(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("sql.ErrNoRows error", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, sql.ErrNoRows)

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, 1)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, 1)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      "email",
			Phone:      "phone",
			State:      1,
			CreatedAt:  model.DateTime(timeNow),
			UpdatedAt:  model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, 1)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      "email",
			Phone:      "phone",
			State:      1,
			CreatedAt:  model.DateTime(timeNow),
			UpdatedAt:  model.DateTime(timeNow.Add(1 * time.Second)),
		}, got)
		assert.NoError(t, err)
	})
}
