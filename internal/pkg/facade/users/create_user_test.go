package users

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFacade_CreateUser(t *testing.T) {
	t.Run("internal error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().Insert(fx.ctx, model.User{
			Name: "name",
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
	})

	t.Run("user already exists error while create user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().Insert(fx.ctx, model.User{
			Name: "name",
		}).Return(0, &mysql.MySQLError{
			Number: 1062,
		})

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().Insert(fx.ctx, model.User{
			Name: "name",
		}).Return(1, nil)

		got, err := fx.facade.CreateUser(fx.ctx, model.User{
			Name: "name",
		})
		assert.Equal(t, int32(1), got)
		assert.NoError(t, err)
	})
}
