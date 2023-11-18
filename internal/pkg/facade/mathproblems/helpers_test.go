package mathproblems

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	mathProblemStorage *mocks.MathProblemStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		mathProblemStorage: mocks.NewMathProblemStorage(t),
	}

	fx.facade = New(Config{
		MathProblemStorage: fx.mathProblemStorage,

		TxManager: fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBMathProblemToModelMathProblem(t *testing.T) {
	type args struct {
		mathProblem database.MathProblem
	}
	tests := []struct {
		name string
		args args
		want model.MathProblem
	}{
		{
			name: "tc1",
			args: args{
				mathProblem: database.MathProblem{
					ID:       1,
					FkGameID: 1,
					URL:      "url",
				},
			},
			want: model.MathProblem{
				ID:     1,
				GameID: 1,
				URL:    "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBMathProblemToModelMathProblem(tt.args.mathProblem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBMathProblemToModelMathProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelMathProblemToDBMathProblem(t *testing.T) {
	type args struct {
		mathProblem model.MathProblem
	}
	tests := []struct {
		name string
		args args
		want database.MathProblem
	}{
		{
			name: "tc1",
			args: args{
				mathProblem: model.MathProblem{
					ID:     1,
					GameID: 1,
					URL:    "url",
				},
			},
			want: database.MathProblem{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelMathProblemToDBMathProblem(tt.args.mathProblem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelMathProblemToDBMathProblem() = %v, want %v", got, tt.want)
			}
		})
	}
}
