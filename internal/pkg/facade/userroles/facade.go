package userroles

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db              *tx.Manager
	userRoleStorage storage.UserRoleStorage
}

// Config ...
type Config struct {
	TxManager       *tx.Manager
	UserRoleStorage storage.UserRoleStorage
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:              cfg.TxManager,
		userRoleStorage: cfg.UserRoleStorage,
	}
}
