package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_UpdatePayment(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("internal error while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some errror"))

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), int32(commonpb.Payment_PAYMENT_CERTIFICATE))
		assert.Error(t, err)
	})

	t.Run("error sql.ErrNoRows while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), int32(commonpb.Payment_PAYMENT_CERTIFICATE))
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), int32(commonpb.Payment_PAYMENT_CERTIFICATE))
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("internal error while update game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:    model.DateTime(timeNow.Add(1 * time.Second)),
			Payment: int32(commonpb.Payment_PAYMENT_CASH),
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, model.Game{
			Date:    model.DateTime(timeNow.Add(1 * time.Second)),
			Payment: int32(commonpb.Payment_PAYMENT_CERTIFICATE),
		}).Return(errors.New("some error"))

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), int32(commonpb.Payment_PAYMENT_CERTIFICATE))
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:    model.DateTime(timeNow.Add(1 * time.Second)),
			Payment: int32(commonpb.Payment_PAYMENT_CASH),
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, model.Game{
			Date:    model.DateTime(timeNow.Add(1 * time.Second)),
			Payment: int32(commonpb.Payment_PAYMENT_CERTIFICATE),
		}).Return(nil)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), int32(commonpb.Payment_PAYMENT_CERTIFICATE))
		assert.NoError(t, err)
	})
}
