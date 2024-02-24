package mathproblem

import (
	"context"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/app/service/math_problem/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	mathproblempb "github.com/nikita5637/quiz-registrator-api/pkg/pb/math_problem"
)

type fixture struct {
	ctx context.Context

	mathProblemsFacade *mocks.MathProblemsFacade

	implementation *Implementation
}

func tearUp(t *testing.T) *fixture {
	fx := &fixture{
		ctx: context.Background(),

		mathProblemsFacade: mocks.NewMathProblemsFacade(t),
	}

	fx.implementation = New(Config{
		MathProblemsFacade: fx.mathProblemsFacade,
	})

	t.Cleanup(func() {})

	return fx
}

func Test_convertModelMathProblemToProtoMathProblem(t *testing.T) {
	type args struct {
		mathProblem model.MathProblem
	}
	tests := []struct {
		name string
		args args
		want *mathproblempb.MathProblem
	}{
		{
			name: "tc1",
			args: args{
				mathProblem: model.MathProblem{
					ID:     1,
					GameID: 2,
					URL:    "url",
				},
			},
			want: &mathproblempb.MathProblem{
				Id:     1,
				GameId: 2,
				Url:    "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelMathProblemToProtoMathProblem(tt.args.mathProblem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelMathProblemToProtoMathProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertProtoMathProblemToModelMathProblem(t *testing.T) {
	type args struct {
		mathProblem *mathproblempb.MathProblem
	}
	tests := []struct {
		name string
		args args
		want model.MathProblem
	}{
		{
			name: "tc1",
			args: args{
				mathProblem: &mathproblempb.MathProblem{
					Id:     1,
					GameId: 2,
					Url:    "url",
				},
			},
			want: model.MathProblem{
				ID:     1,
				GameID: 2,
				URL:    "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertProtoMathProblemToModelMathProblem(tt.args.mathProblem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertProtoMathProblemToModelMathProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getErrorDetails(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		args args
		want *errorDetails
	}{
		{
			name: "keys is nil",
			args: args{
				keys: nil,
			},
			want: nil,
		},
		{
			name: "keys is empty",
			args: args{
				keys: []string{},
			},
			want: nil,
		},
		{
			name: "GameID",
			args: args{
				keys: []string{"GameID"},
			},
			want: &errorDetails{
				Reason: reasonInvalidGameID,
				Lexeme: invalidGameIDLexeme,
			},
		},
		{
			name: "URL",
			args: args{
				keys: []string{"URL"},
			},
			want: &errorDetails{
				Reason: reasonInvalidURL,
				Lexeme: invalidURLLexeme,
			},
		},
		{
			name: "not found",
			args: args{
				keys: []string{"not found"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getErrorDetails(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getErrorDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}
