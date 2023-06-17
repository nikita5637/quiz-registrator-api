package icsfiles

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db             *tx.Manager
	icsFileStorage storage.ICSFileStorage
}

// Config ...
type Config struct {
	ICSFileStorage storage.ICSFileStorage
	TxManager      *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:             cfg.TxManager,
		icsFileStorage: cfg.ICSFileStorage,
	}
}
