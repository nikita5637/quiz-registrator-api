package userroles

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListUserRoles(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRoles(mock.Anything).Return([]database.UserRole{}, errors.New("some error"))

		got, err := fx.facade.ListUserRoles(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userRoleStorage.EXPECT().GetUserRoles(mock.Anything).Return([]database.UserRole{
			{
				ID:       1,
				FkUserID: 1,
				Role:     model.RoleAdmin.String(),
			},
			{
				ID:       2,
				FkUserID: 1,
				Role:     model.RoleManagement.String(),
			},
		}, nil)

		got, err := fx.facade.ListUserRoles(fx.ctx)
		assert.Equal(t, []model.UserRole{
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
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
