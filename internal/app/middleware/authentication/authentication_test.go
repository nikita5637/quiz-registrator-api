package authentication

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestMiddleware_Authentication(t *testing.T) {
	t.Run("error. there are not metadata", func(t *testing.T) {
		fx := tearUp(t)

		fn := fx.middleware.Authentication()
		_, err := fn(fx.ctx)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
	})

	t.Run("error. authentication type not found", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
	})

	t.Run("error. authentication type telegram ID. invalid telegram ID", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				telegramClientIDHeader: "invalid telegram ID",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
	})

	t.Run("error. authentication type telegram ID. get user by telegram ID error", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				telegramClientIDHeader: "1",
			},
		))

		fx.usersFacade.EXPECT().GetUserByTelegramID(ctx, int64(1)).Return(model.User{}, errors.New("some error"))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
	})

	t.Run("ok. authentication type telegram ID", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				telegramClientIDHeader: "1",
			},
		))

		fx.usersFacade.EXPECT().GetUserByTelegramID(ctx, int64(1)).Return(model.User{
			ID:         1,
			TelegramID: 1,
			Name:       "name",
		}, nil)

		fn := fx.middleware.Authentication()
		ctx, err := fn(ctx)
		assert.NoError(t, err)

		userFromContext := usersutils.UserFromContext(ctx)
		assert.Equal(t, model.User{
			ID:         1,
			TelegramID: 1,
			Name:       "name",
		}, userFromContext)
	})

	t.Run("error. authentication type service name. empty service and module names", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				moduleNameHeader:  "",
				serviceNameHeader: "",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Unauthenticated, st.Code())
	})

	t.Run("ok. authentication type service name. service name eq \"fetcher\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				serviceNameHeader: "fetcher",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok. authentication type service name. service name eq \"ics-manager\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				serviceNameHeader: "ics-manager",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok. authentication type service name. service name eq \"telegram\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				serviceNameHeader: "telegram",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok. authentication type service name. service name eq \"telegram-reminder\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				serviceNameHeader: "telegram-reminder",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok. authentication type service name. module name eq \"telegram\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				moduleNameHeader: "telegram",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})

	t.Run("ok. authentication type service name. module name eq \"telegram-reminder\"", func(t *testing.T) {
		fx := tearUp(t)

		ctx := metadata.NewIncomingContext(fx.ctx, metadata.New(
			map[string]string{
				moduleNameHeader: "telegram-reminder",
			},
		))

		fn := fx.middleware.Authentication()
		_, err := fn(ctx)
		assert.NoError(t, err)
	})
}
