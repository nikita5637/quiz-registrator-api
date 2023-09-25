package model

import (
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
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
}

// NewGame returns dummy game
func NewGame() Game {
	return Game{
		ExternalID:  maybe.Nothing[int32](),
		Name:        maybe.Nothing[string](),
		PaymentType: maybe.Nothing[string](),
		Payment:     maybe.Nothing[Payment](),
	}
}

// DateTime ...
func (g *Game) DateTime() DateTime {
	if g == nil {
		return DateTime{}
	}

	return g.Date
}

// IsActive ...
func (g *Game) IsActive() bool {
	if g == nil {
		return false
	}

	activeGameLag := config.GetValue("ActiveGameLag").Uint16()
	return time_utils.TimeNow().Before(g.DateTime().AsTime().Add(time.Duration(activeGameLag) * time.Second))
}
