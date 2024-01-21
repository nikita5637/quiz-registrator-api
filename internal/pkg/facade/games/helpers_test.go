package games

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	dbmocks "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock

	gameStorage *dbmocks.GameStorage
	quizLogger  *mocks.QuizLogger

	facade *Facade
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameStorage: dbmocks.NewGameStorage(t),
		quizLogger:  mocks.NewQuizLogger(t),
	}

	fx.facade = New(Config{
		GameStorage: fx.gameStorage,
		TxManager:   fx.db,
		QuizLogger:  fx.quizLogger,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBGameToModelGame(t *testing.T) {
	timeNow := timeutils.TimeNow()

	type args struct {
		game database.Game
	}
	tests := []struct {
		name string
		args args
		want model.Game
	}{
		{
			name: "test case 1",
			args: args{
				game: database.Game{
					ID: 1,
					ExternalID: sql.NullInt64{
						Int64: 2,
						Valid: true,
					},
					LeagueID: 1,
					Type:     1,
					Number:   "1",
					Name: sql.NullString{
						String: "name",
						Valid:  true,
					},
					PlaceID:     4,
					Date:        timeNow,
					Price:       400,
					PaymentType: []byte("cash,card"),
					MaxPlayers:  9,
					Payment: sql.NullInt64{
						Int64: 1,
						Valid: true,
					},
					Registered: true,
				},
			},
			want: model.Game{
				ID:          1,
				ExternalID:  maybe.Just(int32(2)),
				LeagueID:    int32(1),
				Type:        1,
				Number:      "1",
				Name:        maybe.Just("name"),
				PlaceID:     4,
				Date:        model.DateTime(timeNow),
				Price:       400,
				PaymentType: maybe.Just("cash,card"),
				MaxPlayers:  9,
				Payment:     maybe.Just(model.PaymentCash),
				Registered:  true,
				GameLink:    maybe.Nothing[string](),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertDBGameToModelGame(tt.args.game)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_convertModelGameToDBGame(t *testing.T) {
	timeNow := timeutils.TimeNow()
	type args struct {
		game model.Game
	}
	tests := []struct {
		name string
		args args
		want database.Game
	}{
		{
			name: "test case 1",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(2)),
					LeagueID:    1,
					Type:        1,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: maybe.Just("cash,card"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCash),
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: 1,
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow,
				Price:       400,
				PaymentType: []byte("cash,card"),
				MaxPlayers:  9,
				Payment: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				Registered: true,
			},
		},
		{
			name: "test case 2",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Nothing[int32](),
					LeagueID:    1,
					Type:        1,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: maybe.Just("cash,card"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCash),
					Registered:  true,
				},
			},
			want: database.Game{
				ID:       1,
				LeagueID: 1,
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow,
				Price:       400,
				PaymentType: []byte("cash,card"),
				MaxPlayers:  9,
				Payment: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				Registered: true,
			},
		},
		{
			name: "test case 3",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(2)),
					LeagueID:    1,
					Type:        1,
					Number:      "1",
					Name:        maybe.Nothing[string](),
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: maybe.Just("cash,card"),
					MaxPlayers:  9,
					Payment:     maybe.Just(model.PaymentCash),
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID:    1,
				Type:        1,
				Number:      "1",
				PlaceID:     4,
				Date:        timeNow,
				Price:       400,
				PaymentType: []byte("cash,card"),
				MaxPlayers:  9,
				Payment: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				Registered: true,
			},
		},
		{
			name: "test case 4",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  maybe.Just(int32(2)),
					LeagueID:    1,
					Type:        1,
					Number:      "1",
					Name:        maybe.Just("name"),
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: maybe.Just("cash,card"),
					MaxPlayers:  9,
					Payment:     maybe.Nothing[model.Payment](),
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: 1,
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow,
				Price:       400,
				PaymentType: []byte("cash,card"),
				MaxPlayers:  9,
				Registered:  true,
			},
		},
	}
	for _, tt := range tests {
		got := convertModelGameToDBGame(tt.args.game)
		assert.Equal(t, tt.want, got)
	}
}

func Test_gameHasPassed(t *testing.T) {
	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("game date + lag less then timeutils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(3601 * time.Second)
		}

		g := model.Game{
			Date: model.DateTime(gameDate),
		}
		got := gameHasPassed(g)
		assert.True(t, got)
	})

	t.Run("game date + lag is equal timeutils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(3600 * time.Second)
		}

		g := model.Game{
			Date: model.DateTime(gameDate),
		}
		got := gameHasPassed(g)
		assert.False(t, got)
	})

	t.Run("game date + lag is greater than timeutils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(3599 * time.Second)
		}

		g := model.Game{
			Date: model.DateTime(gameDate),
		}
		got := gameHasPassed(g)
		assert.False(t, got)
	})
}

func Test_getGameLink(t *testing.T) {
	type args struct {
		modelGame model.Game
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "tc1",
			args: args{
				modelGame: model.Game{},
			},
			want: "",
		},
		{
			name: "tc2",
			args: args{
				modelGame: model.Game{
					ExternalID: maybe.Nothing[int32](),
					LeagueID:   model.LeagueQuizPlease,
				},
			},
			want: "",
		},
		{
			name: "tc3",
			args: args{
				modelGame: model.Game{
					ExternalID: maybe.Just(int32(66059)),
					LeagueID:   model.LeagueQuizPlease,
				},
			},
			want: "https://spb.quizplease.ru/game-page?id=66059",
		},
		{
			name: "tc4",
			args: args{
				modelGame: model.Game{
					ExternalID: maybe.Just(int32(21281)),
					LeagueID:   model.LeagueSixtySeconds,
				},
			},
			want: "https://club60sec.ru/quizgames/game/21281/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGameLink(tt.args.modelGame); got != tt.want {
				t.Errorf("getGameLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
