package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestNewUserStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewUserStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}

func Test_convertDBUserToModelUser(t *testing.T) {
	timeNow := time_utils.TimeNow()
	type args struct {
		user User
	}
	tests := []struct {
		name string
		args args
		want model.User
	}{
		{
			name: "test case 1",
			args: args{
				user: User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Email: sql.NullString{
						String: "email",
						Valid:  true,
					},
					State: 1,
					CreatedAt: sql.NullTime{
						Time:  timeNow,
						Valid: true,
					},
				},
			},
			want: model.User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Email:      "email",
				State:      1,
				CreatedAt:  model.DateTime(timeNow),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBUserToModelUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBUserToModelUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelUserToDBUser(t *testing.T) {
	type args struct {
		user model.User
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "test case 1",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Email:      "email",
					Phone:      "phone",
					State:      1,
					CreatedAt:  model.DateTime(time_utils.TimeNow()),
					UpdatedAt:  model.DateTime(time_utils.TimeNow()),
				},
			},
			want: User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Email: sql.NullString{
					String: "email",
					Valid:  true,
				},
				Phone: sql.NullString{
					String: "phone",
					Valid:  true,
				},
				State: 1,
			},
		},
		{
			name: "test case 2",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Email:      "email",
					State:      1,
					CreatedAt:  model.DateTime(time_utils.TimeNow()),
					UpdatedAt:  model.DateTime(time_utils.TimeNow()),
				},
			},
			want: User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Email: sql.NullString{
					String: "email",
					Valid:  true,
				},
				State: 1,
			},
		},
		{
			name: "test case 3",
			args: args{
				user: model.User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Phone:      "phone",
					State:      1,
					CreatedAt:  model.DateTime(time_utils.TimeNow()),
					UpdatedAt:  model.DateTime(time_utils.TimeNow()),
				},
			},
			want: User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Phone: sql.NullString{
					String: "phone",
					Valid:  true,
				},
				State: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelUserToDBUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelUserToDBUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
