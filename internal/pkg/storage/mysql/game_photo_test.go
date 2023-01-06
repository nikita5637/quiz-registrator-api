package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestNewGamePhotoStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGamePhotoStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBGamePhotoToModelGamePhoto(t *testing.T) {
	type args struct {
		game GamePhoto
	}
	tests := []struct {
		name string
		args args
		want model.GamePhoto
	}{
		{
			name: "test case 1",
			args: args{
				game: GamePhoto{
					ID:       1,
					FkGameID: 1,
					URL:      "url",
					CreatedAt: sql.NullTime{
						Time:  time_utils.TimeNow(),
						Valid: true,
					},
				},
			},
			want: model.GamePhoto{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBGamePhotoToModelGamePhoto(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBGamePhotoToModelGamePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelGamePhotoToDBGamePhoto(t *testing.T) {
	type args struct {
		game model.GamePhoto
	}
	tests := []struct {
		name string
		args args
		want GamePhoto
	}{
		{
			name: "test case 1",
			args: args{
				game: model.GamePhoto{
					ID:       1,
					FkGameID: 1,
					URL:      "url",
				},
			},
			want: GamePhoto{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGamePhotoToDBGamePhoto(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGamePhotoToDBGamePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}
