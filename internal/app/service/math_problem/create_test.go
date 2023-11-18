package mathproblem

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/mathproblems"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestImplementation_CreateMathProblem(t *testing.T) {
	t.Run("error: bad request", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{
			MathProblem: &mathproblempb.MathProblem{
				GameId: 0,
				Url:    "http://someurl",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, reasonInvalidGameID, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "GameID: cannot be blank.",
		}, errorInfo.Metadata)
	})

	t.Run("error: internal error while creating math problem", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "http://someurl",
		}).Return(model.MathProblem{}, errors.New("some error"))

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{
			MathProblem: &mathproblempb.MathProblem{
				GameId: 1,
				Url:    "http://someurl",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "http://someurl",
		}).Return(model.MathProblem{}, games.ErrGameNotFound)

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{
			MathProblem: &mathproblempb.MathProblem{
				GameId: 1,
				Url:    "http://someurl",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, games.ReasonGameNotFound, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: math problem already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "http://someurl",
		}).Return(model.MathProblem{}, mathproblems.ErrMathProblemAlreadyExists)

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{
			MathProblem: &mathproblempb.MathProblem{
				GameId: 1,
				Url:    "http://someurl",
			},
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.AlreadyExists, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, mathproblems.ReasonMathProblemAlreadyExists, errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "math problem already exists",
		}, errorInfo.Metadata)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.mathProblemsFacade.EXPECT().CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "http://someurl",
		}).Return(model.MathProblem{
			ID:     1,
			GameID: 1,
			URL:    "http://someurl",
		}, nil)

		got, err := fx.implementation.CreateMathProblem(fx.ctx, &mathproblempb.CreateMathProblemRequest{
			MathProblem: &mathproblempb.MathProblem{
				GameId: 1,
				Url:    "http://someurl",
			},
		})
		assert.NotNil(t, got)
		assert.NoError(t, err)
		assert.Equal(t, &mathproblempb.MathProblem{
			Id:     1,
			GameId: 1,
			Url:    "http://someurl",
		}, got)
	})
}

func Test_validateCreatedMathProblem(t *testing.T) {
	type args struct {
		mathProblem model.MathProblem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "gameID is mepty",
			args: args{
				mathProblem: model.MathProblem{
					GameID: 0,
					URL:    "http://someurl",
				},
			},
			wantErr: true,
		},
		{
			name: "gameID lt 0",
			args: args{
				mathProblem: model.MathProblem{
					GameID: -1,
					URL:    "http://someurl",
				},
			},
			wantErr: true,
		},
		{
			name: "URL is empty",
			args: args{
				mathProblem: model.MathProblem{
					GameID: 1,
					URL:    "",
				},
			},
			wantErr: true,
		},
		{
			name: "URL is not url",
			args: args{
				mathProblem: model.MathProblem{
					GameID: 1,
					URL:    "invalid value",
				},
			},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{
				mathProblem: model.MathProblem{
					GameID: 1,
					URL:    "http://someurl",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCreatedMathProblem(tt.args.mathProblem); (err != nil) != tt.wantErr {
				t.Errorf("validateCreatedMathProblem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
