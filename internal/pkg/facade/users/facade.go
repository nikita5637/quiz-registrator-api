package users

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db          *tx.Manager
	userStorage storage.UserStorage
}

// Config ...
type Config struct {
	UserStorage storage.UserStorage
	TxManager   *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:          cfg.TxManager,
		userStorage: cfg.UserStorage,
	}
}
