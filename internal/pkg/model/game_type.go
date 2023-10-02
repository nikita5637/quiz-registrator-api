package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GameType ...
type GameType int32

const (
	// GameTypeClassic ...
	GameTypeClassic GameType = 1
	// GameTypeThematic ...
	GameTypeThematic GameType = 2
	// GameTypeEnglish ...
	GameTypeEnglish GameType = 4
	// GameTypeMoviesAndMusic ...
	GameTypeMoviesAndMusic GameType = 5
	// GameTypeClosed ...
	GameTypeClosed GameType = 6
	// GameTypeThematicMoviesAndMusic ...
	GameTypeThematicMoviesAndMusic GameType = 9
)

// ValidateGameType ...
func ValidateGameType(value interface{}) error {
	v, ok := value.(GameType)
	if !ok {
		return errors.New("must be GameType")
	}

	return validation.Validate(v, validation.Required, validation.In(
		GameTypeClassic,
		GameTypeThematic,
		GameTypeEnglish,
		GameTypeMoviesAndMusic,
		GameTypeClosed,
		GameTypeThematicMoviesAndMusic,
	))
}
