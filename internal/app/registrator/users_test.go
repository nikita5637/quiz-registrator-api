package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_CreateUser(t *testing.T) {
	t.Run("internal error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateWelcome,
		}).Return(0, errors.New("some error"))

		got, err := fx.registrator.CreateUser(fx.ctx, &registrator.CreateUserRequest{
			Name:       "name",
			TelegramId: -100,
			State:      registrator.UserState(model.UserStateWelcome),
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().CreateUser(fx.ctx, model.User{
			Name:       "name",
			TelegramID: -100,
			State:      model.UserStateWelcome,
		}).Return(0, model.ErrUserAlreadyExists)

		got, err := fx.registrator.CreateUser(fx.ctx, &registrator.CreateUserRequest{
			Name:       "name",
			TelegramId: -100,
			State:      registrator.UserState(model.UserStateWelcome),
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
			State:      model.UserStateWelcome,
		}).Return(1, nil)

		got, err := fx.registrator.CreateUser(fx.ctx, &registrator.CreateUserRequest{
			Name:       "name",
			TelegramId: -100,
			State:      registrator.UserState(model.UserStateWelcome),
		})
		assert.NotNil(t, got)
		assert.Equal(t, int32(1), got.GetId())
		assert.NoError(t, err)
	})
}

func TestRegistrator_GetUser(t *testing.T) {
	t.Run("internal error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx).Return(model.User{}, errors.New("some error"))

		got, err := fx.registrator.GetUser(fx.ctx, &registrator.GetUserRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		user := model.User{
			ID:    1,
			Name:  "name",
			Email: "email",
			Phone: "phone",
			State: model.UserStateChangingName,
		}

		user.TelegramID = 1

		fx.usersFacade.EXPECT().GetUser(fx.ctx).Return(user, nil)

		got, err := fx.registrator.GetUser(fx.ctx, &registrator.GetUserRequest{})
		assert.Equal(t, &registrator.GetUserResponse{
			User: &registrator.User{
				Id:         1,
				Name:       "name",
				Email:      "email",
				Phone:      "phone",
				State:      registrator.UserState(model.UserStateChangingName),
				TelegramId: 1,
			},
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_GetUserByTelegramID(t *testing.T) {
	t.Run("internal error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.registrator.GetUserByTelegramID(fx.ctx, &registrator.GetUserByTelegramIDRequest{
			TelegramId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user not found error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, model.ErrUserNotFound)

		got, err := fx.registrator.GetUserByTelegramID(fx.ctx, &registrator.GetUserByTelegramIDRequest{
			TelegramId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		user := model.User{
			ID:    1,
			Name:  "name",
			Email: "email",
			Phone: "phone",
			State: model.UserStateChangingName,
		}

		user.TelegramID = 1

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(user, nil)

		got, err := fx.registrator.GetUserByTelegramID(fx.ctx, &registrator.GetUserByTelegramIDRequest{
			TelegramId: 1,
		})
		assert.Equal(t, &registrator.GetUserByTelegramIDResponse{
			User: &registrator.User{
				Id:         1,
				Name:       "name",
				Email:      "email",
				Phone:      "phone",
				State:      registrator.UserState(model.UserStateChangingName),
				TelegramId: 1,
			},
		}, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_UpdateUserEmail(t *testing.T) {
	t.Run("\"user not found\" error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserEmail(fx.ctx, int32(1), "some@mail.ru").Return(model.ErrUserNotFound)

		got, err := fx.registrator.UpdateUserEmail(fx.ctx, &registrator.UpdateUserEmailRequest{
			UserId: 1,
			Email:  "some@mail.ru",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.NotFound)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserEmail(fx.ctx, int32(1), "some@mail.").Return(model.ErrUserEmailValidate)

		got, err := fx.registrator.UpdateUserEmail(fx.ctx, &registrator.UpdateUserEmailRequest{
			UserId: 1,
			Email:  "some@mail.",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user email: some@mail.", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errEmailValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserEmail(fx.ctx, int32(1), "some@mail.ru").Return(nil)

		got, err := fx.registrator.UpdateUserEmail(fx.ctx, &registrator.UpdateUserEmailRequest{
			UserId: 1,
			Email:  "some@mail.ru",
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_UpdateUserName(t *testing.T) {
	t.Run("\"user not found\" error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserName(fx.ctx, int32(1), "newName").Return(model.ErrUserNotFound)

		got, err := fx.registrator.UpdateUserName(fx.ctx, &registrator.UpdateUserNameRequest{
			UserId: 1,
			Name:   "newName",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.NotFound)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserName(fx.ctx, int32(1), "newName123").Return(model.ErrUserNameValidateAlphabet)

		got, err := fx.registrator.UpdateUserName(fx.ctx, &registrator.UpdateUserNameRequest{
			UserId: 1,
			Name:   "newName123",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user name: newName123", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errNameAlphabetValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("validation error #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserName(fx.ctx, int32(1), "").Return(model.ErrUserNameValidateLength)

		got, err := fx.registrator.UpdateUserName(fx.ctx, &registrator.UpdateUserNameRequest{
			UserId: 1,
			Name:   "",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user name: ", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errNameLengthValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("validation error #3", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserName(fx.ctx, int32(1), "").Return(model.ErrUserNameValidateRequired)

		got, err := fx.registrator.UpdateUserName(fx.ctx, &registrator.UpdateUserNameRequest{
			UserId: 1,
			Name:   "",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user name: ", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errNameRequiredhValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserName(fx.ctx, int32(1), "newName").Return(nil)

		got, err := fx.registrator.UpdateUserName(fx.ctx, &registrator.UpdateUserNameRequest{
			UserId: 1,
			Name:   "newName",
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_UpdateUserPhone(t *testing.T) {
	t.Run("\"user not found\" error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserPhone(fx.ctx, int32(1), "+79998887766").Return(model.ErrUserNotFound)

		got, err := fx.registrator.UpdateUserPhone(fx.ctx, &registrator.UpdateUserPhoneRequest{
			UserId: 1,
			Phone:  "+79998887766",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.NotFound)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserPhone(fx.ctx, int32(1), "+7").Return(model.ErrUserPhoneValidate)

		got, err := fx.registrator.UpdateUserPhone(fx.ctx, &registrator.UpdateUserPhoneRequest{
			UserId: 1,
			Phone:  "+7",
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user phone: +7", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errPhoneValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserPhone(fx.ctx, int32(1), "+79998887766").Return(nil)

		got, err := fx.registrator.UpdateUserPhone(fx.ctx, &registrator.UpdateUserPhoneRequest{
			UserId: 1,
			Phone:  "+79998887766",
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}

func TestRegistrator_UpdateUserState(t *testing.T) {
	t.Run("\"user not found\" error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserState(fx.ctx, int32(1), model.UserStateRegistered).Return(model.ErrUserNotFound)

		got, err := fx.registrator.UpdateUserState(fx.ctx, &registrator.UpdateUserStateRequest{
			UserId: 1,
			State:  registrator.UserState(model.UserStateRegistered),
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.NotFound)
		assert.Len(t, st.Details(), 2)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserState(fx.ctx, int32(1), model.UserStateInvalid).Return(model.ErrUserStateValidate)

		got, err := fx.registrator.UpdateUserState(fx.ctx, &registrator.UpdateUserStateRequest{
			UserId: 1,
			State:  0,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, st.Code(), codes.InvalidArgument)
		assert.Len(t, st.Details(), 2)

		for _, detail := range st.Details() {
			switch dt := detail.(type) {
			case *errdetails.ErrorInfo:
				assert.Equal(t, "invalid user state: 0", dt.GetReason())
			case *errdetails.LocalizedMessage:
				assert.Equal(t, getTranslator(errStateValidateLexeme)(fx.ctx), dt.GetMessage())
			}
		}
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().UpdateUserState(fx.ctx, int32(1), model.UserStateRegistered).Return(nil)

		got, err := fx.registrator.UpdateUserState(fx.ctx, &registrator.UpdateUserStateRequest{
			UserId: 1,
			State:  2,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
	})
}
