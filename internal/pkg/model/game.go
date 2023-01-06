package model

import (
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// Game ...
type Game struct {
	ID          int32
	ExternalID  int32
	LeagueID    int32
	Type        int32
	Number      string
	Name        string
	PlaceID     int32
	Date        DateTime
	Price       uint32
	PaymentType string
	MaxPlayers  uint32
	Payment     int32
	Registered  bool
	CreatedAt   DateTime
	UpdatedAt   DateTime
	DeletedAt   DateTime
	//
	gameAdditionalInfo
}

type gameAdditionalInfo struct {
	My                  bool
	NumberOfMyLegioners uint32
	NumberOfLegioners   uint32
	NumberOfPlayers     uint32
	ResultPlace         uint32
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
	return g.DeletedAt.AsTime().IsZero() && time_utils.TimeNow().Before(g.DateTime().AsTime().Add(time.Duration(activeGameLag)*time.Second))
}
