package userroles

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateUserRole(t *testing.T) {
	t.Run("error. get user roles error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRolesByUserID(mock.Anything, int(1)).Return(nil, errors.New("some error"))

		got, err := fx.facade.CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		})
		assert.Equal(t, model.UserRole{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. user role already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRolesByUserID(mock.Anything, int(1)).Return([]database.UserRole{
			{
				ID:       1,
				FkUserID: 1,
				Role:     database.Role(model.RoleAdmin),
			},
			{
				ID:       2,
				FkUserID: 1,
				Role:     database.Role(model.RoleUser),
			},
		}, nil)

		got, err := fx.facade.CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		})
		assert.Equal(t, model.UserRole{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserRoleAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. user not found while create user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRolesByUserID(mock.Anything, int(1)).Return([]database.UserRole{}, nil)
		fx.userRoleStorage.EXPECT().Insert(mock.Anything, database.UserRole{
			FkUserID: 1,
			Role:     database.Role(model.RoleAdmin),
		}).Return(0, &mysql.MySQLError{
			Number: 1452,
		})

		got, err := fx.facade.CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		})
		assert.Equal(t, model.UserRole{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. internal error while create user role", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userRoleStorage.EXPECT().GetUserRolesByUserID(mock.Anything, int(1)).Return([]database.UserRole{}, nil)
		fx.userRoleStorage.EXPECT().Insert(mock.Anything, database.UserRole{
			FkUserID: 1,
			Role:     database.Role(model.RoleAdmin),
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		})
		assert.Equal(t, model.UserRole{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userRoleStorage.EXPECT().GetUserRolesByUserID(mock.Anything, int(1)).Return([]database.UserRole{}, nil)
		fx.userRoleStorage.EXPECT().Insert(mock.Anything, database.UserRole{
			FkUserID: 1,
			Role:     database.Role(model.RoleAdmin),
		}).Return(1, nil)

		got, err := fx.facade.CreateUserRole(fx.ctx, model.UserRole{
			UserID: 1,
			Role:   model.RoleAdmin,
		})
		assert.Equal(t, model.UserRole{
			ID:     1,
			UserID: 1,
			Role:   model.RoleAdmin,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
