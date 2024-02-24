package mathproblem

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/mathproblems"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SearchMathProblemByGameID ...
func (i *Implementation) SearchMathProblemByGameID(ctx context.Context, req *mathproblempb.SearchMathProblemByGameIDRequest) (*mathproblempb.MathProblem, error) {
	mathProblem, err := i.mathProblemsFacade.GetMathProblemByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, mathproblems.ErrMathProblemNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, mathproblems.ErrMathProblemNotFound.Error(), mathproblems.ReasonMathProblemNotFound, map[string]string{
				"error": err.Error(),
			}, mathproblems.MathProblemNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelMathProblemToProtoMathProblem(mathProblem), nil
}
