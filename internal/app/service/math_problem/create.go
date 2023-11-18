package mathproblem

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/mathproblems"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMathProblem ...
func (i *Implementation) CreateMathProblem(ctx context.Context, req *mathproblempb.CreateMathProblemRequest) (*mathproblempb.MathProblem, error) {
	if req.GetMathProblem() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	createdMathProblem := convertProtoMathProblemToModelMathProblem(req.GetMathProblem())
	if err := validateCreatedMathProblem(createdMathProblem); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if validationErrors, ok := err.(validation.Errors); ok && len(validationErrors) > 0 {
			keys := make([]string, 0, len(validationErrors))
			for k := range validationErrors {
				keys = append(keys, k)
			}

			if errorDetails := getErrorDetails(keys); errorDetails != nil {
				st = model.GetStatus(ctx,
					codes.InvalidArgument,
					fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()),
					errorDetails.Reason,
					map[string]string{
						"error": err.Error(),
					},
					errorDetails.Lexeme,
				)
			}
		}

		return nil, st.Err()
	}

	mathProblem, err := i.mathProblemsFacade.CreateMathProblem(ctx, createdMathProblem)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		} else if errors.Is(err, mathproblems.ErrMathProblemAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, mathproblems.ErrMathProblemAlreadyExists.Error(), mathproblems.ReasonMathProblemAlreadyExists, map[string]string{
				"error": err.Error(),
			}, mathproblems.MathProblemAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelMathProblemToProtoMathProblem(mathProblem), nil
}

func validateCreatedMathProblem(mathProblem model.MathProblem) error {
	return validation.ValidateStruct(&mathProblem,
		validation.Field(&mathProblem.GameID, validation.Required, validation.Min(1)),
		validation.Field(&mathProblem.URL, validation.Required, is.URL),
	)
}
