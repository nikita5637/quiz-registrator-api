package usermanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	users "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestImplementation_GetUser(t *testing.T) {
	t.Run("internal error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.GetUser(fx.ctx, &usermanagerpb.GetUserRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("user not found error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(model.User{}, users.ErrUserNotFound)

		got, err := fx.implementation.GetUser(fx.ctx, &usermanagerpb.GetUserRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonUserNotFound, errorInfo.Reason)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_MALE

		user := model.User{
			ID:         1,
			Name:       "name",
			TelegramID: 1,
			Email:      maybe.Just("email"),
			Phone:      maybe.Just("phone"),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexMale),
		}

		fx.usersFacade.EXPECT().GetUser(fx.ctx, int32(1)).Return(user, nil)

		got, err := fx.implementation.GetUser(fx.ctx, &usermanagerpb.GetUserRequest{
			Id: 1,
		})
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "name",
			TelegramId: 1,
			Email:      wrapperspb.String("email"),
			Phone:      wrapperspb.String("phone"),
			State:      usermanagerpb.UserState(model.UserStateChangingName),
			Birthdate:  wrapperspb.String("1990-01-30"),
			Sex:        &pbSex,
		}, got)
		assert.NoError(t, err)
	})
}

func TestImplementation_GetUserByTelegramID(t *testing.T) {
	t.Run("internal error while get user", func(t *testing.T) {
		fx := tearUp(t)

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, errors.New("some error"))

		got, err := fx.implementation.GetUserByTelegramID(fx.ctx, &usermanagerpb.GetUserByTelegramIDRequest{
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

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(model.User{}, users.ErrUserNotFound)

		got, err := fx.implementation.GetUserByTelegramID(fx.ctx, &usermanagerpb.GetUserByTelegramIDRequest{
			TelegramId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonUserNotFound, errorInfo.Reason)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		pbSex := usermanagerpb.Sex_SEX_FEMALE

		user := model.User{
			ID:         1,
			Name:       "name",
			TelegramID: 1,
			Email:      maybe.Just("email"),
			Phone:      maybe.Just("phone"),
			State:      model.UserStateChangingName,
			Birthdate:  maybe.Just("1990-01-30"),
			Sex:        maybe.Just(model.SexFemale),
		}

		fx.usersFacade.EXPECT().GetUserByTelegramID(fx.ctx, int64(1)).Return(user, nil)

		got, err := fx.implementation.GetUserByTelegramID(fx.ctx, &usermanagerpb.GetUserByTelegramIDRequest{
			TelegramId: 1,
		})
		assert.Equal(t, &usermanagerpb.User{
			Id:         1,
			Name:       "name",
			TelegramId: 1,
			Email:      wrapperspb.String("email"),
			Phone:      wrapperspb.String("phone"),
			State:      usermanagerpb.UserState(model.UserStateChangingName),
			Birthdate:  wrapperspb.String("1990-01-30"),
			Sex:        &pbSex,
		}, got)
		assert.NoError(t, err)
	})
}
