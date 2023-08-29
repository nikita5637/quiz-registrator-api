package gamephotos

import (
	"context"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	gameStorage      *mocks.GameStorage
	gamePhotoStorage *mocks.GamePhotoStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameStorage:      mocks.NewGameStorage(t),
		gamePhotoStorage: mocks.NewGamePhotoStorage(t),
	}

	fx.facade = &Facade{
		db: fx.db,

		gameStorage:      fx.gameStorage,
		gamePhotoStorage: fx.gamePhotoStorage,
	}

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}
