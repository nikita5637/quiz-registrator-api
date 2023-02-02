package reminder

// Game is a game remind body struct
type Game struct {
	GameID    int32   `json:"game_id,omitempty"`
	PlayerIDs []int32 `json:"player_ids,omitempty"`
}

// Lottery is a lottery remind body struct
type Lottery struct {
	GameID    int32   `json:"game_id,omitempty"`
	PlayerIDs []int32 `json:"player_ids,omitempty"`
}
