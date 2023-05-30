package certificates

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListCertificates(t *testing.T) {
	t.Run("error while getting certificates", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificates(mock.Anything).Return(nil, errors.New("some error"))

		got, err := fx.facade.ListCertificates(fx.ctx)

		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().GetCertificates(mock.Anything).Return([]database.Certificate{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		got, err := fx.facade.ListCertificates(fx.ctx)

		assert.ElementsMatch(t, []model.Certificate{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
