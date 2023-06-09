package i18n

// Lexeme ...
type Lexeme struct {
	Key      string
	FallBack string
}

// Global lexemes ...
var (
	// GameNotFoundLexeme ...
	GameNotFoundLexeme = Lexeme{
		Key:      "game_not_found",
		FallBack: "Game not found",
	}
)
