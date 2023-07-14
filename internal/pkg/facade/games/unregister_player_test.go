package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_UnregisterPlayer(t *testing.T) {
	timeNow := time_utils.TimeNow()

	t.Run("internal error while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 1, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("error sql.ErrNoRows while get game by ID", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 1, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 1, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusInvalid, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("internal error while get game players", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().Or(
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(1),
					"fk_user_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
			),
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
				builder.IsNull{
					"fk_user_id",
				},
			),
		)).Return([]database.GamePlayer{}, errors.New("some error"))

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 1, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("game players not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().Or(
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(1),
					"fk_user_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
			),
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
				builder.IsNull{
					"fk_user_id",
				},
			),
		)).Return([]database.GamePlayer{}, nil)

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 1, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusNotRegistered, got)
		assert.NoError(t, err)
	})

	t.Run("internal error while delete game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().Or(
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"fk_user_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
			),
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
				builder.IsNull{
					"fk_user_id",
				},
			),
		)).Return([]database.GamePlayer{
			{
				ID:       2,
				FkGameID: 2,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().Delete(fx.ctx, int32(2)).Return(errors.New("some error"))

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 2, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusInvalid, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().Or(
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"fk_user_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
			),
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
				builder.IsNull{
					"fk_user_id",
				},
			),
		)).Return([]database.GamePlayer{
			{
				ID:       2,
				FkGameID: 2,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().Delete(fx.ctx, int32(2)).Return(nil)

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 2, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusOK, got)
		assert.NoError(t, err)
	})

	t.Run("not registered yet", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			Date: model.DateTime(timeNow.Add(1 * time.Second)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().Or(
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"fk_user_id":    int32(1),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
			),
			builder.NewCond().And(
				builder.Eq{
					"fk_game_id":    int32(2),
					"registered_by": int32(1),
				},
				builder.IsNull{
					"deleted_at",
				},
				builder.IsNull{
					"fk_user_id",
				},
			),
		)).Return([]database.GamePlayer{
			{
				ID:       3,
				FkGameID: 3,
				FkUserID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
				RegisteredBy: 1,
			},
		}, nil)

		got, err := fx.facade.UnregisterPlayer(fx.ctx, 2, 1, 1)
		assert.Equal(t, model.UnregisterPlayerStatusNotRegistered, got)
		assert.NoError(t, err)
	})
}
