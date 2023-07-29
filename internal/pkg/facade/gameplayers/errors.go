package gameplayers

import "errors"

const (
	// ReasonGamePlayerAlreadyRegistered ...
	ReasonGamePlayerAlreadyRegistered = "GAME_PLAYER_ALREADY_REGISTERED"
	// ReasonGamePlayerNotFound ...
	ReasonGamePlayerNotFound = "GAME_PLAYER_NOT_FOUND"
)

var (
	// ErrGamePlayerAlreadyRegistered ...
	ErrGamePlayerAlreadyRegistered = errors.New("game player already registered")
	// ErrGamePlayerNotFound ...
	ErrGamePlayerNotFound = errors.New("game player not found")
)
