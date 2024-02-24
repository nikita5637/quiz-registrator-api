package certificates

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates/mocks"
	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	dbmocks "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	certificateStorage *dbmocks.CertificateStorage
	quizLogger         mocks.QuizLogger
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		certificateStorage: dbmocks.NewCertificateStorage(t),
		quizLogger:         *mocks.NewQuizLogger(t),
	}

	fx.facade = New(Config{
		CertificateStorage: fx.certificateStorage,
		TxManager:          fx.db,
		QuizLogger:         &fx.quizLogger,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBCertificateToModelCertificate(t *testing.T) {
	type args struct {
		certificate database.Certificate
	}
	tests := []struct {
		name string
		args args
		want model.Certificate
	}{
		{
			name: "tc1",
			args: args{
				certificate: database.Certificate{
					ID:    1,
					Type:  1,
					WonOn: 1,
					SpentOn: sql.NullInt64{
						Int64: 0,
						Valid: false,
					},
					Info: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
			},
			want: model.Certificate{
				ID:      1,
				Type:    1,
				WonOn:   1,
				SpentOn: maybe.Nothing[int32](),
				Info:    maybe.Nothing[string](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBCertificateToModelCertificate(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBCertificateToModelCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelCertificateToDBCertificate(t *testing.T) {
	type args struct {
		certificate model.Certificate
	}
	tests := []struct {
		name string
		args args
		want database.Certificate
	}{
		{
			name: "tc1",
			args: args{
				certificate: model.Certificate{
					ID:      1,
					Type:    model.CertificateTypeFreePass,
					WonOn:   1,
					SpentOn: maybe.Nothing[int32](),
					Info:    maybe.Nothing[string](),
				},
			},
			want: database.Certificate{
				ID:    1,
				Type:  1,
				WonOn: 1,
				SpentOn: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				Info: sql.NullString{
					String: "",
					Valid:  false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelCertificateToDBCertificate(tt.args.certificate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelCertificateToDBCertificate() = %v, want %v", got, tt.want)
			}
		})
	}
}
