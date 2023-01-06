package mysql

import (
	"database/sql"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestNewGameStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGameStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBGameToModelGame(t *testing.T) {
	timeNow := time_utils.TimeNow()

	type args struct {
		game Game
	}
	tests := []struct {
		name string
		args args
		want model.Game
	}{
		{
			name: "test case 1",
			args: args{
				game: Game{
					ID: 1,
					ExternalID: sql.NullInt64{
						Int64: 2,
						Valid: true,
					},
					LeagueID: int(model.LeagueQuizPlease),
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
				LeagueID:    model.LeagueQuizPlease,
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
		want Game
	}{
		{
			name: "test case 1",
			args: args{
				game: model.Game{
					ID:          1,
					ExternalID:  2,
					LeagueID:    model.LeagueQuizPlease,
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
			want: Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: int(model.LeagueQuizPlease),
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
					LeagueID:    model.LeagueQuizPlease,
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
			want: Game{
				ID:       1,
				LeagueID: int(model.LeagueQuizPlease),
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
					LeagueID:    model.LeagueQuizPlease,
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
			want: Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID:    int(model.LeagueQuizPlease),
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
					LeagueID:    model.LeagueQuizPlease,
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
			want: Game{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				LeagueID: int(model.LeagueQuizPlease),
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
