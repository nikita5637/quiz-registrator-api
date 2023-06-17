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

func TestFacade_DeleteICSFile(t *testing.T) {
	t.Run("error. ICS file not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(&database.IcsFile{}, sql.ErrNoRows)

		err := fx.facade.DeleteICSFile(fx.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrICSFileNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("some error while getting ICS file", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(&database.IcsFile{}, errors.New("some error"))

		err := fx.facade.DeleteICSFile(fx.ctx, 1)

		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while deleting ICS file", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.icsFileStorage.EXPECT().GetICSFileByID(mock.Anything, 1).Return(&database.IcsFile{
			ID: 1,
		}, nil)

		fx.icsFileStorage.EXPECT().DeleteICSFile(mock.Anything, 1).Return(errors.New("some error"))

		err := fx.facade.DeleteICSFile(fx.ctx, 1)

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

		fx.icsFileStorage.EXPECT().DeleteICSFile(mock.Anything, 1).Return(nil)

		err := fx.facade.DeleteICSFile(fx.ctx, 1)

		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
