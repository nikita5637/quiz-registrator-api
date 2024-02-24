package certificates

import (
	"database/sql"
	"errors"
	"testing"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_DeleteCertificate(t *testing.T) {
	t.Run("error. certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{}, sql.ErrNoRows)

		err := fx.facade.DeleteCertificate(fx.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrCertificateNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("some error while getting certificate", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{}, errors.New("some error"))

		err := fx.facade.DeleteCertificate(fx.ctx, 1)

		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("try to delete deleted certificate", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{
			ID: 1,
			DeletedAt: sql.NullTime{
				Valid: true,
				Time:  timeutils.TimeNow(),
			},
		}, nil)

		err := fx.facade.DeleteCertificate(fx.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrCertificateNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while deleting certificate", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{
			ID: 1,
		}, nil)

		fx.certificateStorage.EXPECT().DeleteCertificate(mock.Anything, 1).Return(errors.New("some error"))

		err := fx.facade.DeleteCertificate(fx.ctx, 1)

		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{
			ID: 1,
		}, nil)

		fx.certificateStorage.EXPECT().DeleteCertificate(mock.Anything, 1).Return(nil)

		err := fx.facade.DeleteCertificate(fx.ctx, 1)

		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
