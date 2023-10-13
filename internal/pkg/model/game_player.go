package model

import (
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	maybejson "github.com/mono83/maybe/json"
)

// Degree ...
type Degree int32

const (
	// DegreeLikely ...
	DegreeLikely Degree = iota + 1
	// DegreeUnlikely ...
	DegreeUnlikely

	numberOfDegrees
)

// ValidateDegree ...
func ValidateDegree(value interface{}) error {
	v, ok := value.(Degree)
	if !ok {
		return errors.New("must be Degree")
	}

	return validation.Validate(v, validation.Required, validation.Min(DegreeLikely), validation.Max(numberOfDegrees-1))
}

// GamePlayer ...
type GamePlayer struct {
	ID           int32
	GameID       int32
	UserID       maybe.Maybe[int32]
	RegisteredBy int32
	Degree       Degree
}

// MarshalJSON ...
func (gp GamePlayer) MarshalJSON() ([]byte, error) {
	type wrapperGamePlayer struct {
		ID           int32
		GameID       int32
		UserID       maybejson.Maybe[int32]
		RegisteredBy int32
		Degree       Degree
	}

	wgp := wrapperGamePlayer{
		ID:           gp.ID,
		GameID:       gp.GameID,
		UserID:       maybejson.Wrap(gp.UserID),
		RegisteredBy: gp.RegisteredBy,
		Degree:       gp.Degree,
	}
	return json.Marshal(wgp)
}
