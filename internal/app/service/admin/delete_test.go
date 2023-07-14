package admin

import (
	"errors"
	"testing"

	userroles "github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/userroles"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/admin"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_DeleteUserRole(t *testing.T) {
	t.Run("error. role not found while delete user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().DeleteUserRole(fx.ctx, int32(1)).Return(userroles.ErrUserRoleNotFound)

		got, err := fx.implementation.DeleteUserRole(fx.ctx, &admin.DeleteUserRoleRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("error. internal error while delete user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().DeleteUserRole(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.implementation.DeleteUserRole(fx.ctx, &admin.DeleteUserRoleRequest{
			Id: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userRolesFacade.EXPECT().DeleteUserRole(fx.ctx, int32(1)).Return(nil)

		got, err := fx.implementation.DeleteUserRole(fx.ctx, &admin.DeleteUserRoleRequest{
			Id: 1,
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
