package games

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"

var (
	// GameAlreadyExistsLexeme ...
	GameAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "game_already_exists",
		FallBack: "Game already exists",
	}
	// GameHasPassedLexeme ...
	GameHasPassedLexeme = i18n.Lexeme{
		Key:      "game_has_passed",
		FallBack: "Game has passed",
	}
	// GameNotFoundLexeme ...
	GameNotFoundLexeme = i18n.Lexeme{
		Key:      "game_not_found",
		FallBack: "Game not found",
	}
)
