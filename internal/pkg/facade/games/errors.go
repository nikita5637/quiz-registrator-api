package games

import "errors"

const (
	// ReasonGameAlreadyExists ...
	ReasonGameAlreadyExists = "GAME_ALREADY_EXISTS"
	// ReasonGameHasPassed ...
	ReasonGameHasPassed = "GAME_HAS_PASSED"
	// ReasonGameNotFound ...
	ReasonGameNotFound = "GAME_NOT_FOUND"
)

var (
	// ErrGameAlreadyExists ...
	ErrGameAlreadyExists = errors.New("game already exists")
	// ErrGameHasPassed ...
	ErrGameHasPassed = errors.New("game has passed")
	// ErrGameNotFound ...
	ErrGameNotFound = errors.New("game not found")
)
