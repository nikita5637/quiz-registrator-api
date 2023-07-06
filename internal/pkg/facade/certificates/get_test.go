package certificates

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetCertificate(t *testing.T) {
	t.Run("error certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{}, sql.ErrNoRows)

		got, err := fx.facade.GetCertificate(fx.ctx, 1)
		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrCertificateNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while get certificate by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{}, errors.New("some error"))

		got, err := fx.facade.GetCertificate(fx.ctx, 1)
		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&database.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 1,
		}, nil)

		got, err := fx.facade.GetCertificate(fx.ctx, 1)
		assert.Equal(t, model.Certificate{
			ID:    1,
			Type:  model.CertificateTypeFreePass,
			WonOn: 1,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
