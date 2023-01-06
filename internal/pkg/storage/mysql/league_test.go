package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewLeagueStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewLeagueStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBLeagueToModelLeague(t *testing.T) {
	type args struct {
		league League
	}
	tests := []struct {
		name string
		args args
		want model.League
	}{
		{
			name: "test case 1",
			args: args{
				league: League{
					ID:   1,
					Name: "name",
					ShortName: sql.NullString{
						String: "short name",
						Valid:  true,
					},
					LogoLink: sql.NullString{
						String: "logo link",
						Valid:  true,
					},
					WebSite: sql.NullString{
						String: "web site",
						Valid:  true,
					},
				},
			},
			want: model.League{
				ID:        1,
				Name:      "name",
				ShortName: "short name",
				LogoLink:  "logo link",
				WebSite:   "web site",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBLeagueToModelLeague(tt.args.league); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
