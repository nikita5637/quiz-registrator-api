package icsfiles

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateICSFile(t *testing.T) {
	t.Run("error. game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().CreateICSFile(mock.Anything, database.IcsFile{
			FkGameID: 1,
			Name:     "some uuid",
		}).Return(0, &mysql.MySQLError{
			Number: 1452,
		})

		got, err := fx.facade.CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		})

		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. ICS file name already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().CreateICSFile(mock.Anything, database.IcsFile{
			FkGameID: 1,
			Name:     "some uuid",
		}).Return(0, &mysql.MySQLError{
			Number:  1062,
			Message: "Duplicate entry 'some uuid' for key 'name'",
		})

		got, err := fx.facade.CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		})

		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. ICS file for game already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().CreateICSFile(mock.Anything, database.IcsFile{
			FkGameID: 1,
			Name:     "some uuid",
		}).Return(0, &mysql.MySQLError{
			Number: 1062,
		})

		got, err := fx.facade.CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		})

		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("some internal error while creating ICS file", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().CreateICSFile(mock.Anything, database.IcsFile{
			FkGameID: 1,
			Name:     "some uuid",
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		})

		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.icsFileStorage.EXPECT().CreateICSFile(mock.Anything, database.IcsFile{
			FkGameID: 1,
			Name:     "some uuid",
		}).Return(1, nil)

		got, err := fx.facade.CreateICSFile(fx.ctx, model.ICSFile{
			GameID: 1,
			Name:   "some uuid",
		})

		assert.Equal(t, model.ICSFile{
			ID:     1,
			GameID: 1,
			Name:   "some uuid",
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
