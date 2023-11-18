package mathproblem

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/mathproblems"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_SearchMathProblemByGameID(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().GetMathProblemByGameID(fx.ctx, int32(1)).Return(model.MathProblem{}, errors.New("some error"))

		got, err := fx.implementation.SearchMathProblemByGameID(fx.ctx, &mathproblempb.SearchMathProblemByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: math problem not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().GetMathProblemByGameID(fx.ctx, int32(1)).Return(model.MathProblem{}, mathproblems.ErrMathProblemNotFound)

		got, err := fx.implementation.SearchMathProblemByGameID(fx.ctx, &mathproblempb.SearchMathProblemByGameIDRequest{
			GameId: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, mathproblems.ReasonMathProblemNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "math problem not found",
		}, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().GetMathProblemByGameID(fx.ctx, int32(1)).Return(model.MathProblem{
			ID:     1,
			GameID: 1,
			URL:    "url",
		}, nil)

		got, err := fx.implementation.SearchMathProblemByGameID(fx.ctx, &mathproblempb.SearchMathProblemByGameIDRequest{
			GameId: 1,
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
		assert.Equal(t, &mathproblempb.MathProblem{
			Id:     1,
			GameId: 1,
			Url:    "url",
		}, got)
	})
}
