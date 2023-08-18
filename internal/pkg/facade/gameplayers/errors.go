package gameplayers

import "errors"

const (
	// ReasonGamePlayerAlreadyExists ...
	ReasonGamePlayerAlreadyExists = "GAME_PLAYER_ALREADY_EXISTS"
	// ReasonGamePlayerNotFound ...
	ReasonGamePlayerNotFound = "GAME_PLAYER_NOT_FOUND"
)

var (
	// ErrGamePlayerAlreadyExists ...
	ErrGamePlayerAlreadyExists = errors.New("game player already exists")
	// ErrGamePlayerNotFound ...
	ErrGamePlayerNotFound = errors.New("game player not found")
)
