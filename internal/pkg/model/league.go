package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	// LeagueInvalid ...
	LeagueInvalid int32 = iota
	// LeagueQuizPlease ...
	LeagueQuizPlease
	// LeagueSquiz ...
	LeagueSquiz
	// LeagueSixtySeconds ...
	LeagueSixtySeconds
	// LeagueShaker ...
	LeagueShaker
	// LeagueWOW ...
	LeagueWOW
	// LeagueMozgoboynya ...
	LeagueMozgoboynya
	// LeagueNashQuiz ...
	LeagueNashQuiz
	// LeagueSmuzi ...
	LeagueSmuzi
	// LeagueQuizPeace ...
	LeagueQuizPeace

	numberOfLeagues
)

// ValidateLeague ...
func ValidateLeague(value interface{}) error {
	v, ok := value.(int32)
	if !ok {
		return errors.New("must be int32")
	}

	return validation.Validate(v, validation.Max(int(numberOfLeagues-1)))
}

// League ...
type League struct {
	ID        int32
	Name      string
	ShortName string
	LogoLink  string
	WebSite   string
}
