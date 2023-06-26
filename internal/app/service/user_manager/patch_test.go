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
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestImplementation_PatchUser(t *testing.T) {
	t.Run("validation error. invalid user name alphabet", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "name",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:   1,
				Name: "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "invalid email",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "invalid phone",
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      100,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("internal error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, []string{
			"name",
			"telegram_id",
		}).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user not found error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, []string{
			"name",
			"telegram_id",
		}).Return(model.User{}, users.ErrUserNotFound)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("user with specified Telegram ID error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, []string{
			"name",
			"telegram_id",
		}).Return(model.User{}, users.ErrUserTelegramIDAlreadyExists)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("user with specified email error while patch user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, []string{
			"name",
			"telegram_id",
		}).Return(model.User{}, users.ErrUserEmailAlreadyExists)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
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

		fx.usersFacade.EXPECT().PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, []string{
			"name",
			"telegram_id",
		}).Return(model.User{
			ID:         1,
			Name:       "Имя",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, nil)

		got, err := fx.implementation.PatchUser(fx.ctx, &usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Id:         1,
				Name:       "Имя",
				TelegramId: -100,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_WELCOME,
			},
			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"name",
					"telegram_id",
				},
			},
		})
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "Имя",
			TelegramId: -100,
			Email:      "email@email.ru",
			Phone:      "+79998887766",
			State:      usermanagerpb.UserState_USER_STATE_WELCOME,
		}, got)
		assert.NoError(t, err)
	})
}

func Test_validatePatchUserRequest(t *testing.T) {
	t.Run("empty name", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				TelegramId: 111,
			},
		})
		assert.NoError(t, err)
	})

	t.Run("name with invalid characters", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Name: "invalid characters",
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errUserNameAlphabet, err)
	})

	t.Run("invalid name length", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Name: "абвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгдабвгда",
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errUserNameLength, err)
	})

	t.Run("Telegram ID eq 0", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 0,
			},
		})
		assert.NoError(t, err)
	})

	t.Run("invalid email", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
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
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
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
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      usermanagerpb.UserState_USER_STATE_INVALID,
			},
		})
		assert.NoError(t, err)
	})

	t.Run("user state eq 100", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
			User: &usermanagerpb.User{
				Name:       "Имя",
				TelegramId: 111,
				Email:      "email@email.ru",
				Phone:      "+79998887766",
				State:      100,
			},
		})
		assert.Error(t, err)
		assert.Equal(t, errInvalidUserState, err)
	})

	t.Run("ok", func(t *testing.T) {
		err := validatePatchUserRequest(&usermanagerpb.PatchUserRequest{
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
