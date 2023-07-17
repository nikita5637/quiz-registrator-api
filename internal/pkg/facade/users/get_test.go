package users

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetUser(t *testing.T) {
	t.Run("error sql.ErrNoRows", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.GetUser(fx.ctx, 1)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while get user by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetUser(fx.ctx, 1)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(&database.User{
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
			Birthdate: sql.NullTime{
				Time:  birthDateTime,
				Valid: true,
			},
			Sex: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, nil)

		got, err := fx.facade.GetUser(fx.ctx, 1)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email"),
			Phone:      model.NewMaybeString("phone"),
			State:      1,
			Birthdate:  model.NewMaybeString("1990-01-30"),
			Sex:        model.NewMaybeInt32(int32(model.SexMale)),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_GetUserByTelegramID(t *testing.T) {
	t.Run("error sql.ErrNoRows", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByTelegramID(mock.Anything, int64(-100)).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, -100)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while get user by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByTelegramID(mock.Anything, int64(-100)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, -100)
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		birthDateTime, err := time.Parse("2006-01-02", birthDate)
		assert.NoError(t, err)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userStorage.EXPECT().GetUserByTelegramID(mock.Anything, int64(-100)).Return(&database.User{
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
			Birthdate: sql.NullTime{
				Time:  birthDateTime,
				Valid: true,
			},
			Sex: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, nil)

		got, err := fx.facade.GetUserByTelegramID(fx.ctx, -100)
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "name",
			TelegramID: -100,
			Email:      model.NewMaybeString("email"),
			Phone:      model.NewMaybeString("phone"),
			State:      1,
			Birthdate:  model.NewMaybeString("1990-01-30"),
			Sex:        model.NewMaybeInt32(int32(model.SexMale)),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
