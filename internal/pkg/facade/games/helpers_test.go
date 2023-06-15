package games

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games/mocks"
	dbmocks "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	gameStorage       *dbmocks.GameStorage
	gamePlayerStorage *dbmocks.GamePlayerStorage

	rabbitMQProducer *mocks.RabbitMQProducer
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameStorage:       dbmocks.NewGameStorage(t),
		gamePlayerStorage: dbmocks.NewGamePlayerStorage(t),

		rabbitMQProducer: mocks.NewRabbitMQProducer(t),
	}

	fx.facade = NewFacade(Config{
		GameStorage:       fx.gameStorage,
		GamePlayerStorage: fx.gamePlayerStorage,

		RabbitMQProducer: fx.rabbitMQProducer,

		TxManager: fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}
