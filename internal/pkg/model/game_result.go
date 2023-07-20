package model

import "github.com/mono83/maybe"

// GameResult ...
type GameResult struct {
	ID          int32
	FkGameID    int32
	ResultPlace uint32
	RoundPoints maybe.Maybe[string]
}
