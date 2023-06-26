package authorization

import (
	"context"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func okHandler(ctx context.Context, req interface{}) (interface{}, error) {
	return nil, nil
}

func TestMiddleware_Authorization(t *testing.T) {
	t.Run("error, user is banned", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			State: model.UserStateBanned,
		})

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{}, okHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.PermissionDenied, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. grpc rule not found", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateRegistered,
		})

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(ctx, int32(1)).Return([]model.UserRole{
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleAdmin,
			},
			{
				ID:     2,
				UserID: 1,
				Role:   model.RoleManagement,
			},
		}, nil)

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{
			FullMethod: "",
		}, okHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.PermissionDenied, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. internal error while get user roles", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateRegistered,
		})

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(ctx, int32(1)).Return(nil, errors.New("some error"))

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{
			FullMethod: "/registrator.RegistratorService/GetRegisteredGames",
		}, okHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok. public role exists", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateRegistered,
		})

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(ctx, int32(1)).Return([]model.UserRole{
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleAdmin,
			},
			{
				ID:     2,
				UserID: 1,
				Role:   model.RoleManagement,
			},
		}, nil)

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{
			FullMethod: "/registrator.RegistratorService/GetRegisteredGames",
		}, okHandler)
		assert.Nil(t, got)
		assert.NoError(t, err)
	})

	t.Run("error. public role doesn't exists. role not found", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateRegistered,
		})

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(ctx, int32(1)).Return([]model.UserRole{
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleManagement,
			},
		}, nil)

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{
			FullMethod: "/registrator.RegistratorService/UpdatePayment",
		}, okHandler)
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.PermissionDenied, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok. public role doesn't exists", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateRegistered,
		})

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(ctx, int32(1)).Return([]model.UserRole{
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleUser,
			},
			{
				ID:     2,
				UserID: 1,
				Role:   model.RoleManagement,
			},
		}, nil)

		fn := fx.middleware.Authorization()
		got, err := fn(ctx, nil, &grpc.UnaryServerInfo{
			FullMethod: "/registrator.RegistratorService/UpdatePayment",
		}, okHandler)
		assert.Nil(t, got)
		assert.NoError(t, err)
	})
}
