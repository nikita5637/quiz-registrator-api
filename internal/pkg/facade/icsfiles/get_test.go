package icsfiles

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetICSFile(t *testing.T) {
	t.Run("error. ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.GetICSFile(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrICSFileNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. get ICS file by ID other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetICSFile(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(&database.IcsFile{
			ID: 1,
		}, nil)

		got, err := fx.facade.GetICSFile(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{
			ID: 1,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_GetICSFileByGameID(t *testing.T) {
	t.Run("error. ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByGameID(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.GetICSFileByGameID(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrICSFileNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. get ICS file by game ID other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByGameID(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetICSFileByGameID(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.icsFileStorage.EXPECT().GetICSFileByGameID(mock.Anything, 1).Return(&database.IcsFile{
			ID: 1,
		}, nil)

		got, err := fx.facade.GetICSFileByGameID(fx.ctx, 1)
		assert.Equal(t, model.ICSFile{
			ID: 1,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
