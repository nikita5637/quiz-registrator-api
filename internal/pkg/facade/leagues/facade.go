package leagues

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"

// Facade ...
type Facade struct {
	leagueStorage storage.LeagueStorage
}

// Config ...
type Config struct {
	LeagueStorage storage.LeagueStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		leagueStorage: cfg.LeagueStorage,
	}
}
