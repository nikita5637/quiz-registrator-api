package gameresults

import (
	"context"
	"database/sql"
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

	gameResultStorage *mocks.GameResultStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameResultStorage: mocks.NewGameResultStorage(t),
	}

	fx.facade = NewFacade(Config{
		GameResultStorage: fx.gameResultStorage,

		TxManager: fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBGameResultToModelGameResult(t *testing.T) {
	type args struct {
		dbGameResult database.GameResult
	}
	tests := []struct {
		name string
		args args
		want model.GameResult
	}{
		{
			name: "tc1",
			args: args{
				dbGameResult: database.GameResult{
					ID:       1,
					FkGameID: 2,
					Place:    3,
					Points: sql.NullString{
						Valid:  true,
						String: "{}",
					},
				},
			},
			want: model.GameResult{
				ID:          1,
				FkGameID:    2,
				ResultPlace: 3,
				RoundPoints: model.NewMaybeString("{}"),
			},
		},
		{
			name: "tc2",
			args: args{
				dbGameResult: database.GameResult{
					ID:       1,
					FkGameID: 2,
					Place:    3,
					Points: sql.NullString{
						Valid:  false,
						String: "",
					},
				},
			},
			want: model.GameResult{
				ID:          1,
				FkGameID:    2,
				ResultPlace: 3,
				RoundPoints: model.MaybeString{
					Valid: false,
					Value: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBGameResultToModelGameResult(tt.args.dbGameResult); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBGameResultToModelGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelGameResultToDBGameResult(t *testing.T) {
	type args struct {
		modelGameResult model.GameResult
	}
	tests := []struct {
		name string
		args args
		want database.GameResult
	}{
		{
			name: "tc1",
			args: args{
				modelGameResult: model.GameResult{
					ID:          1,
					FkGameID:    2,
					ResultPlace: 3,
					RoundPoints: model.NewMaybeString("{}"),
				},
			},
			want: database.GameResult{
				ID:       1,
				FkGameID: 2,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			},
		},
		{
			name: "tc2",
			args: args{
				modelGameResult: model.GameResult{
					ID:          1,
					FkGameID:    2,
					ResultPlace: 3,
					RoundPoints: model.MaybeString{
						Valid: false,
						Value: "",
					},
				},
			},
			want: database.GameResult{
				ID:       1,
				FkGameID: 2,
				Place:    3,
				Points: sql.NullString{
					Valid:  false,
					String: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGameResultToDBGameResult(tt.args.modelGameResult); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGameResultToDBGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
