package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetGameByID(t *testing.T) {
	// TODO
}

func TestFacade_GetGames(t *testing.T) {
	// TODO tests
}

func TestFacade_GetGamesByUserID(t *testing.T) {
	// TODO tests
}

func TestFacade_GetPlayersByGameID(t *testing.T) {
	t.Run("internal error while getting games", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.facade.GetPlayersByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error sql.ErrNoRows while getting games", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		got, err := fx.facade.GetPlayersByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("found not active game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, nil)

		got, err := fx.facade.GetPlayersByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("error while getting players", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:   1,
			Date: model.DateTime(time_utils.TimeNow().Add(24 * time.Hour)),
		}, nil)

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetPlayersByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:   1,
			Date: model.DateTime(time_utils.TimeNow().Add(24 * time.Hour)),
		}, nil)

		players := []database.GamePlayer{
			{
				ID: 1,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       uint8(registrator.Degree_DEGREE_LIKELY),
			},
			{
				ID: 2,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       uint8(registrator.Degree_DEGREE_UNLIKELY),
			},
			{
				ID: 3,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 3,
				Degree:       uint8(registrator.Degree_DEGREE_UNLIKELY),
			},
		}

		fx.gamePlayerStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(players, nil)

		got, err := fx.facade.GetPlayersByGameID(fx.ctx, 1)

		expected := []model.GamePlayer{
			{
				ID:           1,
				FkUserID:     model.NewMaybeInt32(0),
				RegisteredBy: 1,
				Degree:       int32(registrator.Degree_DEGREE_LIKELY),
			},
			{
				ID:           2,
				FkUserID:     model.NewMaybeInt32(0),
				RegisteredBy: 1,
				Degree:       int32(registrator.Degree_DEGREE_UNLIKELY),
			},
			{
				ID:           3,
				FkUserID:     model.NewMaybeInt32(0),
				RegisteredBy: 3,
				Degree:       int32(registrator.Degree_DEGREE_UNLIKELY),
			},
		}

		assert.Equal(t, expected, got)
		assert.NoError(t, err)
	})
}

func TestFacade_GetRegisteredGames(t *testing.T) {
	// TODO tests
}

func TestFacade_GetTodaysGames(t *testing.T) {
	t.Run("internal error while find games", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}

		fx := tearUp(t)
		fx.gameStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2022-02-10%\""),
		), "").Return(nil, errors.New("some error"))

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
	t.Run("ok", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}

		fx := tearUp(t)
		fx.gameStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2022-02-10%\""),
		), "").Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Len(t, got, 2)
		assert.NoError(t, err)
	})
}
