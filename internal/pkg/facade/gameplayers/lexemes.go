package gameplayers

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"

var (
	// GamePlayerAlreadyExistsLexeme ...
	GamePlayerAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "game_player_already_exists",
		FallBack: "Game player already exists",
	}
	// GamePlayerNotFoundLexeme ...
	GamePlayerNotFoundLexeme = i18n.Lexeme{
		Key:      "game_player_not_found",
		FallBack: "Game player not found",
	}
)
