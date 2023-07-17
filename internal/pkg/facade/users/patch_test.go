package users

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchUser(t *testing.T) {
	t.Run("user with specified telegram ID already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().PatchUser(mock.Anything, database.User{
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
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user with specified email already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().PatchUser(mock.Anything, database.User{
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
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while update", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.userStorage.EXPECT().PatchUser(mock.Anything, database.User{
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
		})
		assert.Equal(t, model.User{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.userStorage.EXPECT().PatchUser(mock.Anything, database.User{
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
		})
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
