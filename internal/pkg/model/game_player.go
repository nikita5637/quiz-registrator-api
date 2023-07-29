package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
)

// Degree ...
type Degree int32

const (
	// DegreeInvalid ...
	DegreeInvalid Degree = iota
	// DegreeLikely ...
	DegreeLikely
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
