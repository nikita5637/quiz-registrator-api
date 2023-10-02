package gameresults

import "errors"

const (
	// ReasonGameResultAlreadyExists ...
	ReasonGameResultAlreadyExists = "GAME_RESULT_ALREADY_EXISTS"
	// ReasonGameResultNotFound ...
	ReasonGameResultNotFound = "GAME_RESULT_NOT_FOUND"
)

var (
	// ErrGameResultAlreadyExists ...
	ErrGameResultAlreadyExists = errors.New("game result already exists")
	// ErrGameResultNotFound ...
	ErrGameResultNotFound = errors.New("game result not found")
)
