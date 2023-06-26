package users

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	usermanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/user_manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchUser(t *testing.T) {
	t.Run("error sql.ErrNoRows while get user by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
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

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user with specified telegram ID already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(&database.User{
			ID:         1,
			Name:       "old name",
			TelegramID: -100,
			Email: sql.NullString{
				String: "old email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "old phone",
				Valid:  true,
			},
			State: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(mock.Anything, database.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email: sql.NullString{
				String: "new email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "new phone",
				Valid:  true,
			},
			State: 2,
		}).Return(&mysql.MySQLError{
			Number:  1062,
			Message: "for key 'telegram_id'",
		})

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user with specified email already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(&database.User{
			ID:         1,
			Name:       "old name",
			TelegramID: -100,
			Email: sql.NullString{
				String: "old email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "old phone",
				Valid:  true,
			},
			State: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(mock.Anything, database.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email: sql.NullString{
				String: "new email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "new phone",
				Valid:  true,
			},
			State: 2,
		}).Return(&mysql.MySQLError{
			Number:  1062,
			Message: "for key 'email'",
		})

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while update", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(&database.User{
			ID:         1,
			Name:       "old name",
			TelegramID: -100,
			Email: sql.NullString{
				String: "old email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "old phone",
				Valid:  true,
			},
			State: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(mock.Anything, database.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email: sql.NullString{
				String: "new email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "new phone",
				Valid:  true,
			},
			State: 2,
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userStorage.EXPECT().GetUserByID(mock.Anything, 1).Return(&database.User{
			ID:         1,
			Name:       "old name",
			TelegramID: -100,
			Email: sql.NullString{
				String: "old email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "old phone",
				Valid:  true,
			},
			State: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(mock.Anything, database.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email: sql.NullString{
				String: "new email",
				Valid:  true,
			},
			Phone: sql.NullString{
				String: "new phone",
				Valid:  true,
			},
			State: 2,
		}).Return(nil)

		got, err := fx.facade.PatchUser(fx.ctx, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, []string{fieldNameName, fieldNameTelegramID, fieldNameEmail, fieldNamePhone, fieldNameState})
		assert.Equal(t, model.User{
			ID:         1,
			Name:       "new name",
			TelegramID: -200,
			Email:      model.NewMaybeString("new email"),
			Phone:      model.NewMaybeString("new phone"),
			State:      model.UserStateRegistered,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_checkPathNames(t *testing.T) {
	field, _ := reflect.ValueOf(usermanagerpb.User{}).Type().FieldByName("Name")
	assert.Equal(t, fieldNameName, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(usermanagerpb.User{}).Type().FieldByName("TelegramId")
	assert.Equal(t, fieldNameTelegramID, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(usermanagerpb.User{}).Type().FieldByName("Email")
	assert.Equal(t, fieldNameEmail, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(usermanagerpb.User{}).Type().FieldByName("Phone")
	assert.Equal(t, fieldNamePhone, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(usermanagerpb.User{}).Type().FieldByName("State")
	assert.Equal(t, fieldNameState, strings.Split(field.Tag.Get("json"), ",")[0])
	assert.Equal(t, 9, reflect.ValueOf(usermanagerpb.User{}).Type().NumField())
}
