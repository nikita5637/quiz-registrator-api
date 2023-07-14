package users

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	userStorage *mocks.UserStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		userStorage: mocks.NewUserStorage(t),
	}

	fx.facade = NewFacade(Config{
		UserStorage: fx.userStorage,
		TxManager:   fx.db,
	})

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBUserToModelUser(t *testing.T) {
	type args struct {
		user database.User
	}
	tests := []struct {
		name string
		args args
		want model.User
	}{
		{
			name: "tc1",
			args: args{
				user: database.User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Email: sql.NullString{
						String: "email@email.ru",
						Valid:  true,
					},
					Phone: sql.NullString{
						String: "+79998887766",
						Valid:  true,
					},
					State: 1,
				},
			},
			want: model.User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Email:      model.NewMaybeString("email@email.ru"),
				Phone:      model.NewMaybeString("+79998887766"),
				State:      model.UserStateWelcome,
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
		want database.User
	}{
		{
			name: "tc1",
			args: args{
				model.User{
					ID:         1,
					Name:       "name",
					TelegramID: -100,
					Email:      model.NewMaybeString("email@email.ru"),
					Phone:      model.NewMaybeString("+79998887766"),
					State:      model.UserStateWelcome,
				},
			},
			want: database.User{
				ID:         1,
				Name:       "name",
				TelegramID: -100,
				Email: sql.NullString{
					String: "email@email.ru",
					Valid:  true,
				},
				Phone: sql.NullString{
					String: "+79998887766",
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
