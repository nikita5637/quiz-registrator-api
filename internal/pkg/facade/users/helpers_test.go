package users

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

const (
	birthDate = "1990-01-30"
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
	t.Run("tc1", func(t *testing.T) {
		got := convertDBUserToModelUser(
			database.User{
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
		)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
		}, got)
	})

	t.Run("tc2", func(t *testing.T) {
		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		got := convertDBUserToModelUser(
			database.User{
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
				Birthdate: sql.NullTime{
					Time:  birthDateTime,
					Valid: true,
				},
			},
		)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate: model.MaybeString{
				Valid: true,
				Value: "1990-01-30",
			},
		}, got)
	})

	t.Run("tc2", func(t *testing.T) {
		got := convertDBUserToModelUser(
			database.User{
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
				Sex: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
		)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Sex:        model.NewMaybeInt32(int32(model.SexMale)),
		}, got)
	})

	t.Run("tc3", func(t *testing.T) {
		got := convertDBUserToModelUser(
			database.User{
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
				Sex: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
			},
		)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Sex:        model.NewMaybeInt32(int32(model.SexFemale)),
		}, got)
	})
}

func Test_convertModelUserToDBUser(t *testing.T) {
	t.Run("tc1", func(t *testing.T) {
		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		got := convertModelUserToDBUser(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  model.NewMaybeString(birthDate),
			Sex:        model.NewMaybeInt32(int32(model.SexMale)),
		})
		assert.Equal(t, database.User{
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
			Birthdate: sql.NullTime{
				Time:  birthDateTime,
				Valid: true,
			},
			Sex: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, got)
	})

	t.Run("tc2", func(t *testing.T) {
		birthDate := "0001-01-01"

		got := convertModelUserToDBUser(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  model.NewMaybeString(birthDate),
			Sex:        model.NewMaybeInt32(int32(model.SexMale)),
		})
		assert.Equal(t, database.User{
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
			Birthdate: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			Sex: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, got)
	})

	t.Run("tc3", func(t *testing.T) {
		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		got := convertModelUserToDBUser(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  model.NewMaybeString(birthDate),
			Sex: model.MaybeInt32{
				Valid: false,
				Value: 0,
			},
		})
		assert.Equal(t, database.User{
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
			Birthdate: sql.NullTime{
				Time:  birthDateTime,
				Valid: true,
			},
			Sex: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		}, got)
	})

	t.Run("tc4", func(t *testing.T) {
		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		got := convertModelUserToDBUser(model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email@email.ru"),
			Phone:      model.NewMaybeString("+79998887766"),
			State:      model.UserStateWelcome,
			Birthdate:  model.NewMaybeString(birthDate),
		})
		assert.Equal(t, database.User{
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
			Birthdate: sql.NullTime{
				Time:  birthDateTime,
				Valid: true,
			},
			Sex: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
		}, got)
	})
}
