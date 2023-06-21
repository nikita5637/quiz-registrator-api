package gameplayers

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db                *tx.Manager
	gamePlayerStorage storage.GamePlayerStorage
}

// Config ...
type Config struct {
	GamePlayerStorage storage.GamePlayerStorage
	TxManager         *tx.Manager
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:                cfg.TxManager,
		gamePlayerStorage: cfg.GamePlayerStorage,
	}
}
