package model

import (
	"github.com/mono83/maybe"
)

// Game ...
type Game struct {
	ID          int32
	ExternalID  maybe.Maybe[int32]
	LeagueID    int32
	Type        GameType
	Number      string
	Name        maybe.Maybe[string]
	PlaceID     int32
	Date        DateTime
	Price       uint32
	PaymentType maybe.Maybe[string]
	MaxPlayers  uint32
	Payment     maybe.Maybe[Payment]
	Registered  bool
	IsInMaster  bool
	// additional info
	HasPassed bool
	GameLink  maybe.Maybe[string]
}

// NewGame returns dummy game
func NewGame() Game {
	return Game{
		ExternalID:  maybe.Nothing[int32](),
		Name:        maybe.Nothing[string](),
		PaymentType: maybe.Nothing[string](),
		Payment:     maybe.Nothing[Payment](),
		GameLink:    maybe.Nothing[string](),
	}
}

// DateTime ...
func (g *Game) DateTime() DateTime {
	if g == nil {
		return DateTime{}
	}

	return g.Date
}
