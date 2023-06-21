package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewGamePlayerStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGamePlayerStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertModelGamePlayerToDBGamePlayer(t *testing.T) {
	type args struct {
		gamePlayer model.GamePlayer
	}
	tests := []struct {
		name string
		args args
		want GamePlayer
	}{
		{
			name: "test case 1",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					FkGameID:     2,
					FkUserID:     model.NewMaybeInt32(3),
					RegisteredBy: 4,
					Degree:       1,
				},
			},
			want: GamePlayer{
				ID:       1,
				FkGameID: 2,
				FkUserID: sql.NullInt64{
					Int64: 3,
					Valid: true,
				},
				RegisteredBy: 4,
				Degree:       1,
			},
		},
		{
			name: "test case 2",
			args: args{
				gamePlayer: model.GamePlayer{
					ID:           1,
					FkGameID:     2,
					RegisteredBy: 4,
					Degree:       1,
				},
			},
			want: GamePlayer{
				ID:           1,
				FkGameID:     2,
				RegisteredBy: 4,
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
