package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewPlaceStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewPlaceStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBPlaceToModelPlace(t *testing.T) {
	type args struct {
		place Place
	}
	tests := []struct {
		name string
		args args
		want model.Place
	}{
		{
			name: "test case 1",
			args: args{
				place: Place{
					ID:      1,
					Name:    "name",
					Address: "address",
					ShortName: sql.NullString{
						String: "short name",
						Valid:  true,
					},
					Latitude: sql.NullFloat64{
						Float64: 1.1,
						Valid:   true,
					},
					Longitude: sql.NullFloat64{
						Float64: 2.2,
						Valid:   true,
					},
					MenuLink: sql.NullString{
						String: "menu link",
						Valid:  true,
					},
				},
			},
			want: model.Place{
				ID:        1,
				Name:      "name",
				Address:   "address",
				ShortName: "short name",
				Latitude:  1.1,
				Longitude: 2.2,
				MenuLink:  "menu link",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBPlaceToModelPlace(tt.args.place); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBPlaceToModelPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
