package model

const (
	// LeagueInvalid ...
	LeagueInvalid int32 = iota
	// LeagueQuizPlease ...
	LeagueQuizPlease
	// LeagueSquiz ...
	LeagueSquiz
	// League60Seconds ...
	League60Seconds
	// LeagueShaker ...
	LeagueShaker

	// Add new leagues here ...

	// NumberOfLeagues ...
	NumberOfLeagues
)

// League ...
type League struct {
	ID        int32
	Name      string
	ShortName string
	LogoLink  string
	WebSite   string
}
