package gameplayers

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	dbmocks "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	gamePlayerStorage *dbmocks.GamePlayerStorage
	quizLogger        *mocks.QuizLogger
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gamePlayerStorage: dbmocks.NewGamePlayerStorage(t),
		quizLogger:        mocks.NewQuizLogger(t),
	}

	fx.facade = New(Config{
		GamePlayerStorage: fx.gamePlayerStorage,
		TxManager:         fx.db,
		QuizLogger:        fx.quizLogger,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBGamePlayerToModelGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer database.GamePlayer
	}
	tests := []struct {
		name string
		args args
		want model.GamePlayer
	}{
		{
			name: "tc1",
			args: args{
				gamePlayer: database.GamePlayer{
					ID:       1,
					FkGameID: 1,
					FkUserID: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
					RegisteredBy: 1,
					Degree:       1,
				},
			},
			want: model.GamePlayer{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Just(int32(1)),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBGamePlayerToModelGamePlayer(tt.args.gamePlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBGamePlayerToModelGamePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelGamePlayerToDBGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer model.GamePlayer
	}
	tests := []struct {
		name string
		args args
		want database.GamePlayer
	}{
		{
			name: "tc1",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					GameID:       1,
					UserID:       maybe.Just(int32(1)),
					RegisteredBy: 1,
					Degree:       model.DegreeLikely,
				},
			},
			want: database.GamePlayer{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
				Degree:       1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGamePlayerToDBGamePlayer(tt.args.gamePlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGamePlayerToDBGamePlayer() = %v, want %v", got, tt.want)
			}
		})
	}
}
