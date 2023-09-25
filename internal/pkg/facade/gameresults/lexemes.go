package gameresults

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"

var (
	// GameResultAlreadyExistsLexeme ...
	GameResultAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "game_result_already_exists",
		FallBack: "Game result already exists",
	}
	// GameResultNotFoundLexeme ...
	GameResultNotFoundLexeme = i18n.Lexeme{
		Key:      "game_result_not_found",
		FallBack: "Game result not found",
	}
)
