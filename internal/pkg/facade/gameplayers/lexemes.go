package gameplayers

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"

var (
	// GamePlayerAlreadyRegisteredLexeme ...
	GamePlayerAlreadyRegisteredLexeme = i18n.Lexeme{
		Key:      "game_player_already_registered",
		FallBack: "Game player already registered",
	}
	// GamePlayerNotFoundLexeme ...
	GamePlayerNotFoundLexeme = i18n.Lexeme{
		Key:      "game_player_not_found",
		FallBack: "Game player not found",
	}
)
