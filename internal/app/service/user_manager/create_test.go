package usermanager

import (
	"errors"
	"testing"

	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_CreateUser(t *testing.T) {
	t.Run("validation error. empty user name", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "",
				TelegramId: -100,
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. user name length gt 100", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "abcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghijabcdefghija",
				TelegramId: -100,
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. empty Telegram ID", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: 0,
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid email format", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "invalid email",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid phone format", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "invalid phone",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error. invalid user state", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState(model.UserStateInvalid),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      model.NewMaybeString("+79998887766"),
			Email:      model.NewMaybeString("email@email.ru"),
			State:      model.UserStateWelcome,
		}).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user with specified Telegram ID already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      model.NewMaybeString("+79998887766"),
			Email:      model.NewMaybeString("email@email.ru"),
			State:      model.UserStateWelcome,
		}).Return(model.User{}, users.ErrUserTelegramIDAlreadyExists)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("user with specified email already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      model.NewMaybeString("+79998887766"),
			Email:      model.NewMaybeString("email@email.ru"),
			State:      model.UserStateWelcome,
		}).Return(model.User{}, users.ErrUserEmailAlreadyExists)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			Phone:      model.NewMaybeString("+79998887766"),
			Email:      model.NewMaybeString("email@email.ru"),
			State:      model.UserStateWelcome,
		}).Return(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Phone:      model.NewMaybeString("+79998887766"),
			Email:      model.NewMaybeString("email@email.ru"),
			State:      model.UserStateWelcome,
		}, nil)

		got, err := fx.implementation.CreateUser(fx.ctx, &usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "name",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState(model.UserStateWelcome),
			},
		})
		assert.NotNil(t, got)
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "name",
			TelegramId: -100,
			Email:      "email@email.ru",
			Phone:      "+79998887766",
			State:      usermanagerpb.UserState_USER_STATE_WELCOME,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validateCreateUserRequest(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				TelegramId: 111,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errUserNameIsRequired, err)
	})

	t.Run("invalid name length", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name: "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errUserNameLength, err)
	})

	t.Run("Telegram ID eq 0", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 0,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errInvalidTelegramID, err)
	})

	t.Run("invalid email", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "invalid email",
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errInvalidEmailFormat, err)
	})

	t.Run("invalid phone", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "email@email.ru",
				Phone:      "invalid phone",
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errInvalidPhoneFormat, err)
	})

	t.Run("invalid user state", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_INVALID,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errInvalidUserState, err)
	})

	t.Run("ok", func(t *testing.T) {
		err := validateCreateUserRequest(&usermanagerpb.CreateUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_CHANGING_NAME,
			},
		})
		assert.NoError(t, err)
	})
}
