package places

import "errors"

const (
	// ReasonPlaceNotFound ...
	ReasonPlaceNotFound = "PLACE_NOT_FOUND"
)

var (
	// ErrPlaceNotFound ...
	ErrPlaceNotFound = errors.New("place not found")
)
