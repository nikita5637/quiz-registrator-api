package mathproblems

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

const (
	mathProblemIBFK1 = "math_problem_ibfk_1"
)

// CreateMathProblem ...
func (f *Facade) CreateMathProblem(ctx context.Context, mathProblem model.MathProblem) (model.MathProblem, error) {
	createdModelMathProblem := model.MathProblem{}
	err := f.db.RunTX(ctx, "CreateMathProblem", func(ctx context.Context) error {
		newDBMathProblem := convertModelMathProblemToDBMathProblem(mathProblem)
		id, err := f.mathProblemStorage.CreateMathProblem(ctx, newDBMathProblem)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, mathProblemIBFK1); i != -1 {
						return fmt.Errorf("create math problem error: %w", games.ErrGameNotFound)
					}
				} else if err.Number == 1062 {
					return ErrMathProblemAlreadyExists
				}
			}

			return fmt.Errorf("create math problem error: %w", err)
		}

		newDBMathProblem.ID = id
		createdModelMathProblem = convertDBMathProblemToModelMathProblem(newDBMathProblem)

		return nil
	})
	if err != nil {
		return model.MathProblem{}, fmt.Errorf("CreateMathProblem error: %w", err)
	}

	return createdModelMathProblem, nil
}
