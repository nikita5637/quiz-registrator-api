package gameplayers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetGamePlayer(t *testing.T) {
	t.Run("game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		got, err := fx.facade.GetGamePlayer(fx.ctx, 1)
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("other error while get game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGamePlayer(fx.ctx, 1)
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("game is deleted", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{
			ID:       1,
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			RegisteredBy: 1,
			Degree:       1,
			DeletedAt: sql.NullTime{
				Time:  time_utils.TimeNow(),
				Valid: true,
			},
		}, nil)

		got, err := fx.facade.GetGamePlayer(fx.ctx, 1)
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{
			ID:       1,
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			RegisteredBy: 1,
			Degree:       1,
		}, nil)

		got, err := fx.facade.GetGamePlayer(fx.ctx, 1)
		assert.Equal(t, model.GamePlayer{
			ID:           1,
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_GetGamePlayersByFields(t *testing.T) {
	t.Run("find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    int32(1),
				"registered_by": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		).And(
			builder.Eq{
				"fk_user_id": int32(1),
			},
		)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok legioner", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    int32(1),
				"registered_by": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		).And(
			builder.IsNull{
				"fk_user_id",
			},
		)).Return([]database.GamePlayer{
			{
				ID:           1,
				FkGameID:     1,
				RegisteredBy: 1,
				Degree:       uint8(model.DegreeLikely),
			},
		}, nil)

		got, err := fx.facade.GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
		})
		assert.Equal(t, []model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok main player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    int32(1),
				"registered_by": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		).And(
			builder.Eq{
				"fk_user_id": int32(1),
			},
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
				Degree:       uint8(model.DegreeLikely),
			},
		}, nil)

		got, err := fx.facade.GetGamePlayersByFields(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
		})
		assert.Equal(t, []model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Just(int32(1)),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_GetGamePlayersByGameID(t *testing.T) {
	t.Run("find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGamePlayersByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
				Degree:       1,
			},
			{
				ID:       2,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       1,
			},
		}, nil)

		got, err := fx.facade.GetGamePlayersByGameID(fx.ctx, 1)
		assert.Equal(t, []model.GamePlayer{
			{
				ID:           1,
				GameID:       1,
				UserID:       maybe.Just(int32(1)),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
			{
				ID:           2,
				GameID:       1,
				UserID:       maybe.Nothing[int32](),
				RegisteredBy: 1,
				Degree:       model.DegreeLikely,
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
