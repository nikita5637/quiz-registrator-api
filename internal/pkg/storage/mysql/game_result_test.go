package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewGameResultStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGameResultStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBGameResultToModelGameResult(t *testing.T) {
	type args struct {
		game GameResult
	}
	tests := []struct {
		name string
		args args
		want model.GameResult
	}{
		{
			name: "test case 1",
			args: args{
				game: GameResult{
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
				RoundPoints: model.MaybeString{
					Valid: true,
					Value: "{}",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBGameResultToModelGameResult(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBGameResultToModelGameResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
