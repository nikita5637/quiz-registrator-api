package model

// GameResult ...
type GameResult struct {
	ID          int32
	FkGameID    int32
	ResultPlace uint32
	RoundPoints MaybeString
}
