//go:generate mockery --case underscore --name MathProblemsFacade --with-expecter

package mathproblem

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
)

// MathProblemsFacade ...
type MathProblemsFacade interface {
	CreateMathProblem(ctx context.Context, mathProblem model.MathProblem) (model.MathProblem, error)
	GetMathProblemByGameID(ctx context.Context, gameID int32) (model.MathProblem, error)
}

// Implementation ...
type Implementation struct {
	mathProblemsFacade MathProblemsFacade

	mathproblempb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	MathProblemsFacade MathProblemsFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		mathProblemsFacade: cfg.MathProblemsFacade,
	}
}
