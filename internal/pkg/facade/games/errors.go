package games

import "errors"

const (
	// ReasonGameNotFound ...
	ReasonGameNotFound = "GAME_NOT_FOUND"
)

var (
	// ErrGameNotFound ...
	ErrGameNotFound = errors.New("game not found")
)
