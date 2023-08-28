package games

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games/mocks"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	dbmocks "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	gameStorage       *dbmocks.GameStorage
	gamePlayerStorage *dbmocks.GamePlayerStorage

	rabbitMQProducer *mocks.RabbitMQProducer
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameStorage:       dbmocks.NewGameStorage(t),
		gamePlayerStorage: dbmocks.NewGamePlayerStorage(t),

		rabbitMQProducer: mocks.NewRabbitMQProducer(t),
	}

	fx.facade = NewFacade(Config{
		GameStorage:       fx.gameStorage,
		GamePlayerStorage: fx.gamePlayerStorage,

		RabbitMQProducer: fx.rabbitMQProducer,

		TxManager: fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBGameToModelGame(t *testing.T) {
	timeNow := time_utils.TimeNow()

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
					LeagueID: int(leaguepb.LeagueID_QUIZ_PLEASE),
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
					CreatedAt: sql.NullTime{
						Time:  timeNow,
						Valid: true,
					},
					UpdatedAt: sql.NullTime{
						Time:  timeNow,
						Valid: true,
					},
					DeletedAt: sql.NullTime{
						Time:  timeNow,
						Valid: true,
					},
				},
			},
			want: model.Game{
				ID:          1,
				ExternalID:  2,
				LeagueID:    int32(leaguepb.LeagueID_QUIZ_PLEASE),
				Type:        1,
				Number:      "1",
				Name:        "name",
				PlaceID:     4,
				Date:        model.DateTime(timeNow),
				Price:       400,
				PaymentType: "cash,card",
				MaxPlayers:  9,
				Payment:     1,
				Registered:  true,
				CreatedAt:   model.DateTime(timeNow),
				UpdatedAt:   model.DateTime(timeNow),
				DeletedAt:   model.DateTime(timeNow),
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
	timeNow := time_utils.TimeNow()
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
					ExternalID:  2,
					LeagueID:    int32(leaguepb.LeagueID_QUIZ_PLEASE),
					Type:        1,
					Number:      "1",
					Name:        "name",
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: "cash,card",
					MaxPlayers:  9,
					Payment:     1,
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: int(leaguepb.LeagueID_QUIZ_PLEASE),
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow.UTC(),
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
					LeagueID:    int32(leaguepb.LeagueID_QUIZ_PLEASE),
					Type:        1,
					Number:      "1",
					Name:        "name",
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: "cash,card",
					MaxPlayers:  9,
					Payment:     1,
					Registered:  true,
				},
			},
			want: database.Game{
				ID:       1,
				LeagueID: int(leaguepb.LeagueID_QUIZ_PLEASE),
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow.UTC(),
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
					ExternalID:  2,
					LeagueID:    int32(leaguepb.LeagueID_QUIZ_PLEASE),
					Type:        1,
					Number:      "1",
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: "cash,card",
					MaxPlayers:  9,
					Payment:     1,
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID:    int(leaguepb.LeagueID_QUIZ_PLEASE),
				Type:        1,
				Number:      "1",
				PlaceID:     4,
				Date:        timeNow.UTC(),
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
					ExternalID:  2,
					LeagueID:    int32(leaguepb.LeagueID_QUIZ_PLEASE),
					Type:        1,
					Number:      "1",
					Name:        "name",
					PlaceID:     4,
					Date:        model.DateTime(timeNow),
					Price:       400,
					PaymentType: "cash,card",
					MaxPlayers:  9,
					Registered:  true,
				},
			},
			want: database.Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: int(leaguepb.LeagueID_QUIZ_PLEASE),
				Type:     1,
				Number:   "1",
				Name: sql.NullString{
					String: "name",
					Valid:  true,
				},
				PlaceID:     4,
				Date:        timeNow.UTC(),
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
