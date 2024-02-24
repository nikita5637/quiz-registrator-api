package mathproblem

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

const (
	reasonInvalidGameID = "INVALID_GAME_ID"
	reasonInvalidURL    = "INVALID_URL"
)

var (
	invalidGameIDLexeme = i18n.Lexeme{
		Key:      "invalid_game_id",
		FallBack: "Invalid game ID",
	}
	invalidURLLexeme = i18n.Lexeme{
		Key:      "invalid_url",
		FallBack: "Invalid URL",
	}

	errorDetailsByField = map[string]errorDetails{
		"GameID": {
			Reason: reasonInvalidGameID,
			Lexeme: invalidGameIDLexeme,
		},
		"URL": {
			Reason: reasonInvalidURL,
			Lexeme: invalidURLLexeme,
		},
	}
)

func convertModelMathProblemToProtoMathProblem(mathProblem model.MathProblem) *mathproblempb.MathProblem {
	return &mathproblempb.MathProblem{
		Id:     mathProblem.ID,
		GameId: mathProblem.GameID,
		Url:    mathProblem.URL,
	}
}

func convertProtoMathProblemToModelMathProblem(mathProblem *mathproblempb.MathProblem) model.MathProblem {
	return model.MathProblem{
		ID:     mathProblem.GetId(),
		GameID: mathProblem.GetGameId(),
		URL:    mathProblem.GetUrl(),
	}
}

func getErrorDetails(keys []string) *errorDetails {
	if len(keys) == 0 {
		return nil
	}

	if v, ok := errorDetailsByField[keys[0]]; ok {
		return &v
	}

	return nil
}
