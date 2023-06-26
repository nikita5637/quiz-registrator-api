package users

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateUser(t *testing.T) {
	t.Run("user with specified Telegram ID already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().Insert(mock.Anything, database.User{
			Name: "name",
		}).Return(0, &mysql.MySQLError{
			Number:  1062,
			Message: "for key 'telegram_id'",
		})

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserTelegramIDAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user with specified email already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().Insert(mock.Anything, database.User{
			Name: "name",
		}).Return(0, &mysql.MySQLError{
			Number:  1062,
			Message: "for key 'email'",
		})

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserEmailAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while insert user", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().Insert(mock.Anything, database.User{
			Name: "name",
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userStorage.EXPECT().Insert(mock.Anything, database.User{
			Name: "name",
		}).Return(1, nil)

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, model.User{
			ID:   1,
			Name: "name",
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
