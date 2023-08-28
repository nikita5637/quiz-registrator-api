package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_UpdatePayment(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("internal error while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(nil, errors.New("some errror"))

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate)
		assert.Error(t, err)
	})

	t.Run("error sql.ErrNoRows while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(nil, sql.ErrNoRows)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{}, nil)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)
	})

	t.Run("internal error while update game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			Date: timeNow.Add(1 * time.Second).UTC(),
			Payment: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, database.Game{
			Date: timeNow.Add(1 * time.Second).UTC(),
			Payment: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
		}).Return(errors.New("some error"))

		err := fx.facade.UpdatePayment(fx.ctx, 1, model.PaymentCertificate)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			Date: timeNow.Add(1 * time.Second).UTC(),
			Payment: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		}, nil)

		fx.gameStorage.EXPECT().Update(fx.ctx, database.Game{
			Date: timeNow.Add(1 * time.Second).UTC(),
			Payment: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
		}).Return(nil)

		err := fx.facade.UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate)
		assert.NoError(t, err)
	})
}
