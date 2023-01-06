package users

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestFacade_UpdateUserEmail(t *testing.T) {
	t.Run("internal error while get user by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "email@mail.ru")
		assert.Error(t, err)
	})

	t.Run("sql.ErrNoRows error while get user by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, sql.ErrNoRows)

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "email@mail.ru")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "email@mail")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserEmailValidate)
	})

	t.Run("empty email", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Email: "",
			State: model.UserStateRegistered,
		}).Return(nil)

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "")
		assert.NoError(t, err)
	})

	t.Run("error while update user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Email: "email@mail.ru",
			State: model.UserStateRegistered,
		}).Return(errors.New("some error"))

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "email@mail.ru")
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Email: "email@mail.ru",
			State: model.UserStateRegistered,
		}).Return(nil)

		err := fx.facade.UpdateUserEmail(fx.ctx, 1, "email@mail.ru")
		assert.NoError(t, err)
	})
}

func TestFacade_UpdateUserName(t *testing.T) {
	t.Run("internal error while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Имя")
		assert.Error(t, err)
	})

	t.Run("sql.ErrNoRows while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, sql.ErrNoRows)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Имя")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("validation error #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNameValidateRequired)
	})

	t.Run("validation error #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Name")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNameValidateAlphabet)
	})

	t.Run("validation error #3", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "ДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмяДлинноеИмя")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNameValidateLength)
	})

	t.Run("error while update user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Name:  "Имя",
			State: model.UserStateRegistered,
		}).Return(errors.New("some error"))

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Имя")
		assert.Error(t, err)
	})

	t.Run("test case 1", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Name:  "Имя Фамилия",
			State: model.UserStateRegistered,
		}).Return(nil)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Имя Фамилия")
		assert.NoError(t, err)
	})

	t.Run("test case 2", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Name:  "Имя",
			State: model.UserStateRegistered,
		}).Return(nil)

		err := fx.facade.UpdateUserName(fx.ctx, 1, "Имя")
		assert.NoError(t, err)
	})
}

func TestFacade_UpdateUserPhone(t *testing.T) {
	t.Run("internal error while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		err := fx.facade.UpdateUserPhone(fx.ctx, 1, "+79998887766")
		assert.Error(t, err)
	})

	t.Run("sql.ErrNoRows error while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, sql.ErrNoRows)

		err := fx.facade.UpdateUserPhone(fx.ctx, 1, "+79998887766")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserPhone(fx.ctx, 1, "+7999888")
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserPhoneValidate)
	})

	t.Run("error while update user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Phone: "+79998887766",
			State: model.UserStateRegistered,
		}).Return(errors.New("some error"))

		err := fx.facade.UpdateUserPhone(fx.ctx, 1, "+79998887766")
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			Phone: "+79998887766",
			State: model.UserStateRegistered,
		}).Return(nil)

		err := fx.facade.UpdateUserPhone(fx.ctx, 1, "+79998887766")
		assert.NoError(t, err)
	})
}

func TestFacade_UpdateUserState(t *testing.T) {
	t.Run("internal error while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, errors.New("some error"))

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStateChangingName)
		assert.Error(t, err)
	})

	t.Run("sql.ErrNoRows while get user by id", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{}, sql.ErrNoRows)

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStateChangingName)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserNotFound)
	})

	t.Run("validation error #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStateInvalid)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserStateValidate)
	})

	t.Run("validation error #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStatesNumber)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrUserStateValidate)
	})

	t.Run("error while update user", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateChangingName,
		}).Return(errors.New("some error"))

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStateChangingName)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.userStorage.EXPECT().GetUserByID(fx.ctx, int32(1)).Return(model.User{
			ID: 1,
		}, nil)

		fx.userStorage.EXPECT().Update(fx.ctx, model.User{
			ID:    1,
			State: model.UserStateChangingName,
		}).Return(nil)

		err := fx.facade.UpdateUserState(fx.ctx, 1, model.UserStateChangingName)
		assert.NoError(t, err)
	})
}
