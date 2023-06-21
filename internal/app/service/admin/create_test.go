package admin

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_CreateUserRole(t *testing.T) {
	t.Run("error. invalid role", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 1,
				Role:   adminpb.Role_ROLE_INVALID,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. user role already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{}, model.ErrUserRoleAlreadyExists)

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 1,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. user not found while create user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{}, model.ErrUserNotFound)

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 1,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. internal error while create user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{}, errors.New("some error"))

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 1,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{
			ID:     1,
			UserID: 1,
			Role:   model.RoleAdmin,
		}, nil)

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 1,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		})
		assert.Equal(t, &adminpb.UserRole{
			Id:     1,
			UserId: 1,
			Role:   adminpb.Role_ROLE_ADMIN,
		}, got)
		assert.NoError(t, err)
	})
}
