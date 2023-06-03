package certificates

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	db "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchCertificate(t *testing.T) {
	t.Run("error while get original certificate. certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{}, sql.ErrNoRows)

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   3,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameWonOn})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrCertificateNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while get original certificate. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{}, errors.New("some error"))

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   3,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameWonOn})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while patch certificate. won_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 3,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 4,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "",
			},
		}, nil)

		fx.certificateStorage.EXPECT().PatchCertificate(mock.Anything, db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: -10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 4,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "",
			},
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: gameIDFK1ConstraintName,
		})

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   -10,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameWonOn})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrWonOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while patch certificate. spent_on game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 4,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "",
			},
		}, nil)

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
				String: "",
			},
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: "",
		})

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(-10),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameSpentOn})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrSpentOnGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while patch certificate. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 10,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 4,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "",
			},
		}, nil)

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
				String: "",
			},
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchCertificate(fx.ctx, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeFreePass,
			WonOn:   10,
			SpentOn: model.NewMaybeInt32(-10),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameSpentOn})

		assert.Equal(t, model.Certificate{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.certificateStorage.EXPECT().GetCertificateByID(mock.Anything, 1).Return(&db.Certificate{
			ID:    1,
			Type:  1,
			WonOn: 3,
			SpentOn: sql.NullInt64{
				Valid: true,
				Int64: 4,
			},
			Info: sql.NullString{
				Valid:  true,
				String: "",
			},
		}, nil)

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
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   1,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, []string{fieldNameType, fieldNameSpentOn, fieldNameInfo})

		assert.Equal(t, model.Certificate{
			ID:      1,
			Type:    pkgmodel.CertificateTypeBarBillPayment,
			WonOn:   3,
			SpentOn: model.NewMaybeInt32(2),
			Info:    model.NewMaybeString("{\"sum\":5000}"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_checkPathNames(t *testing.T) {
	field, _ := reflect.ValueOf(registrator.Certificate{}).Type().FieldByName("Type")
	assert.Equal(t, fieldNameType, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(registrator.Certificate{}).Type().FieldByName("WonOn")
	assert.Equal(t, fieldNameWonOn, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(registrator.Certificate{}).Type().FieldByName("SpentOn")
	assert.Equal(t, fieldNameSpentOn, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(registrator.Certificate{}).Type().FieldByName("Info")
	assert.Equal(t, fieldNameInfo, strings.Split(field.Tag.Get("json"), ",")[0])
	assert.Equal(t, 8, reflect.ValueOf(registrator.Certificate{}).Type().NumField())
}
