package leagues

import "errors"

const (
	// ReasonLeagueNotFound ...
	ReasonLeagueNotFound = "LEAGUE_NOT_FOUND"
)

var (
	// ErrLeagueNotFound ...
	ErrLeagueNotFound = errors.New("league not found")
)
