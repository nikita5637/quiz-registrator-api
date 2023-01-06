package gamephotos

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"

// Facade ...
type Facade struct {
	gameStorage       storage.GameStorage
	gamePhotoStorage  storage.GamePhotoStorage
	gameResultStorage storage.GameResultStorage
}

// Config ...
type Config struct {
	GameStorage       storage.GameStorage
	GamePhotoStorage  storage.GamePhotoStorage
	GameResultStorage storage.GameResultStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		gameStorage:       cfg.GameStorage,
		gamePhotoStorage:  cfg.GamePhotoStorage,
		gameResultStorage: cfg.GameResultStorage,
	}
}
