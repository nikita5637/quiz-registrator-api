package users

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"

// Facade ...
type Facade struct {
	userStorage storage.UserStorage
}

// Config ...
type Config struct {
	UserStorage storage.UserStorage
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		userStorage: cfg.UserStorage,
	}
}
