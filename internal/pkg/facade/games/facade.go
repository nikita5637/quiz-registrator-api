package games

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db                *tx.Manager
	gameStorage       storage.GameStorage
	gamePlayerStorage storage.GamePlayerStorage
}

// Config ...
type Config struct {
	GameStorage       storage.GameStorage
	GamePlayerStorage storage.GamePlayerStorage
	TxManager         *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:                cfg.TxManager,
		gameStorage:       cfg.GameStorage,
		gamePlayerStorage: cfg.GamePlayerStorage,
	}
}
