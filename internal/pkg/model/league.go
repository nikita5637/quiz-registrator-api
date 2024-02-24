package model

const (
	// LeagueQuizPlease ...
	LeagueQuizPlease = int32(1)
	// LeagueSquiz ...
	LeagueSquiz = int32(2)
	// LeagueSixtySeconds ...
	LeagueSixtySeconds = int32(3)
)

// League ...
type League struct {
	ID        int32
	Name      string
	ShortName string
	LogoLink  string
	WebSite   string
}
