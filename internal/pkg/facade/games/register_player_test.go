package games

import (
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_RegisterPlayer(t *testing.T) {
	t.Run("empty fk game ID", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.RegisterPlayer(fx.ctx, 0, 1, 1, int32(registrator.Degree_DEGREE_LIKELY))
		assert.Equal(t, model.RegisterPlayerStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("game not active", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:       model.DateTime(time_utils.TimeNow().Add(-24 * time.Hour)),
			MaxPlayers: 3,
		}, nil)

		got, err := fx.facade.RegisterPlayer(fx.ctx, 1, 1, 1, int32(registrator.Degree_DEGREE_LIKELY))
		assert.Equal(t, model.RegisterPlayerStatus(model.RegisterPlayerStatusInvalid), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("player is registered", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:       model.DateTime(time_utils.TimeNow().Add(1 * time.Hour)),
			MaxPlayers: 3,
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]model.GamePlayer{
			{
				ID:       1,
				FkUserID: 1,
			},
			{
				ID:       2,
				FkUserID: 2,
			},
			{
				ID:       3,
				FkUserID: 3,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
				"fk_user_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]model.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: 1,
			},
		}, nil)

		got, err := fx.facade.RegisterPlayer(fx.ctx, 1, 1, 1, int32(registrator.Degree_DEGREE_LIKELY))
		assert.Equal(t, model.RegisterPlayerStatus(model.RegisterPlayerStatusAlreadyRegistered), got)
		assert.NoError(t, err)
	})

	t.Run("player is not registered", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date:       model.DateTime(time_utils.TimeNow().Add(1 * time.Hour)),
			MaxPlayers: 3,
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]model.GamePlayer{
			{
				ID:       2,
				FkUserID: 2,
			},
			{
				ID:       3,
				FkUserID: 3,
			},
			{
				ID:       4,
				FkUserID: 4,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
				"fk_user_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]model.GamePlayer{}, nil)

		got, err := fx.facade.RegisterPlayer(fx.ctx, 1, 1, 1, int32(registrator.Degree_DEGREE_LIKELY))
		assert.Equal(t, model.RegisterPlayerStatus(model.RegisterPlayerStatusInvalid), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNoFreeSlots)
	})
}
