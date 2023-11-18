package mathproblems

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetMathProblemByGameID ...
func (f *Facade) GetMathProblemByGameID(ctx context.Context, gameID int32) (model.MathProblem, error) {
	mathProblem := model.MathProblem{}
	err := f.db.RunTX(ctx, "GetMathProblemByGameID", func(ctx context.Context) error {
		dbMathProblems, err := f.mathProblemStorage.GetMathProblemByGameID(ctx, int(gameID))
		if err != nil {
			return fmt.Errorf("get math problem by game ID error: %w", err)
		}

		if len(dbMathProblems) == 0 {
			return ErrMathProblemNotFound
		}

		mathProblem = convertDBMathProblemToModelMathProblem(*dbMathProblems[0])

		return nil
	})
	if err != nil {
		return model.MathProblem{}, fmt.Errorf("GetMathProblemByGameID error: %w", err)
	}

	return mathProblem, nil
}
