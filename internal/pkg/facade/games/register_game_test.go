package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_RegisterGame(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("internal error while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("error game not found while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("internal error while update game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Payment:    int32(registrator.Payment_PAYMENT_CASH),
			Registered: true,
		}).Return(errors.New("some error"))

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("game already registered", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: true,
		}, nil)

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusAlreadyRegistered, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Registered: false,
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, model.Game{
			Date:       model.DateTime(timeNow.Add(1 * time.Second)),
			Payment:    int32(registrator.Payment_PAYMENT_CASH),
			Registered: true,
		}).Return(nil)

		got, err := fx.facade.RegisterGame(fx.ctx, 1)
		assert.Equal(t, model.RegisterGameStatusOK, got)
		assert.NoError(t, err)
	})
}
