package places

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"

// Facade ...
type Facade struct {
	placeStorage storage.PlaceStorage
}

// Config ...
type Config struct {
	PlaceStorage storage.PlaceStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		placeStorage: cfg.PlaceStorage,
	}
}
