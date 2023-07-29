package admin

import (
	"errors"
	"testing"

	userroles "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/userroles"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_CreateUserRole(t *testing.T) {
	t.Run("error. invalid user ID", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateUserRole(fx.ctx, &adminpb.CreateUserRoleRequest{
			UserRole: &adminpb.UserRole{
				UserId: 0,
				Role:   adminpb.Role_ROLE_ADMIN,
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserID, errorInfo.Reason)
	})

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
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidUserRole, errorInfo.Reason)
	})

	t.Run("error. user role already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{}, userroles.ErrRoleIsAlreadyAssigned)

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
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonRoleIsAlreadyAssign, errorInfo.Reason)
	})

	t.Run("error. user not found while create user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		}).Return(model.UserRole{}, users.ErrUserNotFound)

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
		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, users.ReasonUserNotFound, errorInfo.Reason)
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

func Test_validateCreatedUserRole(t *testing.T) {
	type args struct {
		userRole model.UserRole
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty user ID",
			args: args{
				userRole: model.UserRole{
					UserID: 0,
					Role:   model.RoleAdmin,
				},
			},
			wantErr: true,
		},
		{
			name: "invalid user role",
			args: args{
				userRole: model.UserRole{
					UserID: 1,
					Role:   model.RoleInvalid,
				},
			},
			wantErr: true,
		},
		{
			name: "user role gt max",
			args: args{
				userRole: model.UserRole{
					UserID: 1,
					Role:   100,
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				userRole: model.UserRole{
					UserID: 1,
					Role:   model.RoleAdmin,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedUserRole(tt.args.userRole); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedUserRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
