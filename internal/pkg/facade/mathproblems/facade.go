package mathproblems

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db                 *tx.Manager
	mathProblemStorage storage.MathProblemStorage
}

// Config ...
type Config struct {
	MathProblemStorage storage.MathProblemStorage
	TxManager          *tx.Manager
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:                 cfg.TxManager,
		mathProblemStorage: cfg.MathProblemStorage,
	}
}
