package games

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"

// Facade ...
type Facade struct {
	gameStorage       storage.GameStorage
	gamePlayerStorage storage.GamePlayerStorage
}

// Config ...
type Config struct {
	GameStorage       storage.GameStorage
	GamePlayerStorage storage.GamePlayerStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		gameStorage:       cfg.GameStorage,
		gamePlayerStorage: cfg.GamePlayerStorage,
	}
}
