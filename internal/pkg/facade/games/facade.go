//go:generate mockery --case underscore --name RabbitMQProducer --with-expecter

package games

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// RabbitMQProducer ...
type RabbitMQProducer interface {
	Send(ctx context.Context, message interface{}) error
}

// Facade ...
type Facade struct {
	db                *tx.Manager
	gameStorage       storage.GameStorage
	gamePlayerStorage storage.GamePlayerStorage
	rabbitMQProducer  RabbitMQProducer
}

// Config ...
type Config struct {
	GameStorage       storage.GameStorage
	GamePlayerStorage storage.GamePlayerStorage
	RabbitMQProducer  RabbitMQProducer
	TxManager         *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:                cfg.TxManager,
		gameStorage:       cfg.GameStorage,
		gamePlayerStorage: cfg.GamePlayerStorage,
		rabbitMQProducer:  cfg.RabbitMQProducer,
	}
}
