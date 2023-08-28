package model

import (
	"errors"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// Game ...
type Game struct {
	ID          int32
	ExternalID  int32
	LeagueID    int32
	Type        GameType
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
	// additional info
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

// ValidateGame ...
func ValidateGame(game Game) error {
	err := validation.Validate(game.LeagueID, validation.Required, validation.Min(1), validation.Max(len(leaguepb.LeagueID_value)-1))
	if err != nil {
		return ErrInvalidLeagueID
	}

	err = validation.Validate(game.Type, validation.Required, validation.By(validateGameType))
	if err != nil {
		return ErrInvalidGameType
	}

	if game.Type != GameTypeClosed {
		err = validation.Validate(game.Number, validation.Required)
		if err != nil {
			return ErrInvalidGameNumber
		}
	}

	err = validation.Validate(game.PlaceID, validation.Required)
	if err != nil {
		return ErrInvalidPlaceID
	}

	err = validation.Validate(game.Date, validation.Required, validation.By(validateGameDate))
	if err != nil {
		return ErrInvalidDate
	}

	err = validation.Validate(game.Price, validation.Required)
	if err != nil {
		return ErrInvalidPrice
	}

	err = validation.Validate(game.MaxPlayers, validation.Required)
	if err != nil {
		return ErrInvalidMaxPlayers
	}

	return nil
}

func validateGameType(value interface{}) error {
	gameType, ok := value.(GameType)
	if !ok {
		return errors.New("must be GameType")
	}

	if gameType == GameTypeClassic ||
		gameType == GameTypeThematic ||
		gameType == GameTypeEnglish ||
		gameType == GameTypeMoviesAndMusic ||
		gameType == GameTypeClosed ||
		gameType == GameTypeThematicMoviesAndMusic {
		return nil
	}

	return ErrInvalidGameType
}

func validateGameDate(value interface{}) error {
	gameDate, ok := value.(DateTime)
	if !ok {
		return errors.New("game date is not model.DateTime")
	}

	if gameDate.AsTime().IsZero() {
		return ErrInvalidDate
	}

	return nil
}
