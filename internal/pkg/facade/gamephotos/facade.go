package gamephotos

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db               *tx.Manager
	gameStorage      storage.GameStorage
	gamePhotoStorage storage.GamePhotoStorage
}

// Config ...
type Config struct {
	GameStorage      storage.GameStorage
	GamePhotoStorage storage.GamePhotoStorage
	TxManager        *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:               cfg.TxManager,
		gameStorage:      cfg.GameStorage,
		gamePhotoStorage: cfg.GamePhotoStorage,
	}
}
