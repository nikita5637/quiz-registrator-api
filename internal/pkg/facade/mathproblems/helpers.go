package mathproblems

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

func convertDBMathProblemToModelMathProblem(mathProblem database.MathProblem) model.MathProblem {
	return model.MathProblem{
		ID:     int32(mathProblem.ID),
		GameID: int32(mathProblem.FkGameID),
		URL:    mathProblem.URL,
	}
}

func convertModelMathProblemToDBMathProblem(mathProblem model.MathProblem) database.MathProblem {
	return database.MathProblem{
		ID:       int(mathProblem.ID),
		FkGameID: int(mathProblem.GameID),
		URL:      mathProblem.URL,
	}
}
