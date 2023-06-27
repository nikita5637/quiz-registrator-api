package certificates

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateCertificate(t *testing.T) {
	t.Run("error. won_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().CreateCertificate(mock.Anything, database.Certificate{
			Type:  2,
			WonOn: 1,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{}",
			},
			CreatedAt: sql.NullTime{},
			UpdatedAt: sql.NullTime{},
			DeletedAt: sql.NullTime{},
		}).Return(0, &mysql.MySQLError{
			Message: gameIDFK1ConstraintName,
			Number:  1452,
		})

		got, err := fx.facade.CreateCertificate(fx.ctx, model.Certificate{
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrWonOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. spent_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().CreateCertificate(mock.Anything, database.Certificate{
			Type:  2,
			WonOn: 1,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{}",
			},
			CreatedAt: sql.NullTime{},
			UpdatedAt: sql.NullTime{},
			DeletedAt: sql.NullTime{},
		}).Return(0, &mysql.MySQLError{
			Message: "",
			Number:  1452,
		})

		got, err := fx.facade.CreateCertificate(fx.ctx, model.Certificate{
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrSpentOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("some internal error while creating certificate", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().CreateCertificate(mock.Anything, database.Certificate{
			Type:  2,
			WonOn: 1,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{}",
			},
			CreatedAt: sql.NullTime{},
			UpdatedAt: sql.NullTime{},
			DeletedAt: sql.NullTime{},
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateCertificate(fx.ctx, model.Certificate{
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{}"),
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

		fx.certificateStorage.EXPECT().CreateCertificate(mock.Anything, database.Certificate{
			Type:  2,
			WonOn: 1,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 2,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "{}",
			},
			CreatedAt: sql.NullTime{},
			UpdatedAt: sql.NullTime{},
			DeletedAt: sql.NullTime{},
		}).Return(1, nil)

		got, err := fx.facade.CreateCertificate(fx.ctx, model.Certificate{
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.Certificate{
			ID:      1,
			Type:    model.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{}"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
