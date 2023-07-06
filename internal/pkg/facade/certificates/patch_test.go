package certificates

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	db "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchCertificate(t *testing.T) {
	t.Run("error while patch certificate. won_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().PatchCertificate(mock.Anything, db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: -10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{\"sum\":5000}",
			},
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: gameIDFK1ConstraintName,
		})

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   -10,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrWonOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while patch certificate. spent_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().PatchCertificate(mock.Anything, db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: -10,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{\"sum\":5000}",
			},
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: "",
		})

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(-10),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrSpentOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while patch certificate. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().PatchCertificate(mock.Anything, db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: -10,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{\"sum\":5000}",
			},
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(-10),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().PatchCertificate(mock.Anything, db.Certificate{
			ID:    1,
			Type:  2,
			WonOn: 3,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{\"sum\":5000}",
			},
		}).Return(nil)

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   3,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		})

		assert.Equal(t, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   3,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
