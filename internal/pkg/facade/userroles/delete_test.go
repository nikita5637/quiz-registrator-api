package userroles

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_DeleteUserRole(t *testing.T) {
	t.Run("error. sql.ErrNoRows", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRoleByID(mock.Anything, 1).Return(&database.UserRole{}, sql.ErrNoRows)

		err := fx.facade.DeleteUserRole(fx.ctx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserRoleNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. internal error while get user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRoleByID(mock.Anything, 1).Return(&database.UserRole{}, errors.New("some error"))

		err := fx.facade.DeleteUserRole(fx.ctx, 1)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. found deleted user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRoleByID(mock.Anything, 1).Return(&database.UserRole{
			ID: 1,
			DeletedAt: sql.NullTime{
				Valid: true,
				Time:  timeutils.TimeNow(),
			},
		}, nil)

		err := fx.facade.DeleteUserRole(fx.ctx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserRoleNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. error while delete user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRoleByID(mock.Anything, 1).Return(&database.UserRole{
			ID: 1,
			DeletedAt: sql.NullTime{
				Valid: false,
				Time:  time.Time{},
			},
		}, nil)

		fx.userRoleStorage.EXPECT().DeleteUserRole(mock.Anything, 1).Return(errors.New("some error"))

		err := fx.facade.DeleteUserRole(fx.ctx, 1)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userRoleStorage.EXPECT().GetUserRoleByID(mock.Anything, 1).Return(&database.UserRole{
			ID: 1,
			DeletedAt: sql.NullTime{
				Valid: false,
				Time:  time.Time{},
			},
		}, nil)

		fx.userRoleStorage.EXPECT().DeleteUserRole(mock.Anything, 1).Return(nil)

		err := fx.facade.DeleteUserRole(fx.ctx, 1)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
