package gameresults

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db                *tx.Manager
	gameResultStorage storage.GameResultStorage
}

// Config ...
type Config struct {
	GameResultStorage storage.GameResultStorage
	TxManager         *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:                cfg.TxManager,
		gameResultStorage: cfg.GameResultStorage,
	}
}
