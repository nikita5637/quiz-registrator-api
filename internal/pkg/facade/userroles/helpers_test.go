package userroles

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

	userRoleStorage *mocks.UserRoleStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		userRoleStorage: mocks.NewUserRoleStorage(t),
	}

	fx.facade = New(Config{
		TxManager:       fx.db,
		UserRoleStorage: fx.userRoleStorage,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}
