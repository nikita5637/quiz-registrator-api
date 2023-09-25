package games

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db          *tx.Manager
	gameStorage storage.GameStorage
}

// Config ...
type Config struct {
	GameStorage storage.GameStorage
	TxManager   *tx.Manager
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:          cfg.TxManager,
		gameStorage: cfg.GameStorage,
	}
}
