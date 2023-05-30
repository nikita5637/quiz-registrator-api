package certificates

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	certificateStorage *mocks.CertificateStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		certificateStorage: mocks.NewCertificateStorage(t),
	}

	fx.facade = NewFacade(Config{
		CertificateStorage: fx.certificateStorage,

		TxManager: fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}
