package adminservice

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	adminpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_GetUserRolesByUserID(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(fx.ctx, int32(1)).Return([]model.UserRole{}, errors.New("some error"))

		got, err := fx.implementation.GetUserRolesByUserID(fx.ctx, &admin.GetUserRolesByUserIDRequest{
			UserId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().GetUserRolesByUserID(fx.ctx, int32(1)).Return([]model.UserRole{
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleAdmin,
			},
			{
				ID:     1,
				UserID: 1,
				Role:   model.RoleManagement,
			},
		}, nil)

		got, err := fx.implementation.GetUserRolesByUserID(fx.ctx, &admin.GetUserRolesByUserIDRequest{
			UserId: 1,
		})
		assert.Equal(t, &adminpb.GetUserRolesByUserIDResponse{
			UserRoles: []*adminpb.UserRole{
				{
					Id:     1,
					UserId: 1,
					Role:   admin.Role_ROLE_ADMIN,
				},
				{
					Id:     1,
					UserId: 1,
					Role:   admin.Role_ROLE_MANAGEMENT,
				},
			},
		}, got)
		assert.NoError(t, err)
	})
}
