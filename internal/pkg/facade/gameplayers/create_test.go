package gameplayers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateGamePlayer(t *testing.T) {
	t.Run("GetGamePlayersByGameID error", func(t *testing.T) {
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

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("game player already registered", func(t *testing.T) {
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
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				RegisteredBy: 1,
				Degree:       2,
			},
		}, nil)

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerAlreadyRegistered)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. user not found", func(t *testing.T) {
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
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       2,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().CreateGamePlayer(mock.Anything, database.GamePlayer{
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			RegisteredBy: 1,
			Degree:       1,
			CreatedAt:    sql.NullTime{},
			UpdatedAt:    sql.NullTime{},
			DeletedAt:    sql.NullTime{},
		}).Return(0, &mysql.MySQLError{
			Message: gameIDFK1ConstraintName,
			Number:  1452,
		})

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, users.ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error. game not found", func(t *testing.T) {
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
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       2,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().CreateGamePlayer(mock.Anything, database.GamePlayer{
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			RegisteredBy: 1,
			Degree:       1,
			CreatedAt:    sql.NullTime{},
			UpdatedAt:    sql.NullTime{},
			DeletedAt:    sql.NullTime{},
		}).Return(0, &mysql.MySQLError{
			Message: gameIDFK2ConstraintName,
			Number:  1452,
		})

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("other error while create game player", func(t *testing.T) {
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
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       2,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().CreateGamePlayer(mock.Anything, database.GamePlayer{
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Valid: true,
				Int64: 1,
			},
			RegisteredBy: 1,
			Degree:       1,
			CreatedAt:    sql.NullTime{},
			UpdatedAt:    sql.NullTime{},
			DeletedAt:    sql.NullTime{},
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok legioner", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().CreateGamePlayer(mock.Anything, database.GamePlayer{
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			RegisteredBy: 1,
			Degree:       1,
			CreatedAt:    sql.NullTime{},
			UpdatedAt:    sql.NullTime{},
			DeletedAt:    sql.NullTime{},
		}).Return(2, nil)

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{
			ID:           2,
			GameID:       1,
			UserID:       maybe.Nothing[int32](),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
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
					Int64: 0,
					Valid: false,
				},
				RegisteredBy: 1,
				Degree:       2,
			},
		}, nil)

		fx.gamePlayerStorage.EXPECT().CreateGamePlayer(mock.Anything, database.GamePlayer{
			FkGameID: 1,
			FkUserID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			RegisteredBy: 1,
			Degree:       1,
			CreatedAt:    sql.NullTime{},
			UpdatedAt:    sql.NullTime{},
			DeletedAt:    sql.NullTime{},
		}).Return(2, nil)

		got, err := fx.facade.CreateGamePlayer(fx.ctx, model.GamePlayer{
			GameID:       1,
			UserID:       maybe.Just(int32(1)),
			RegisteredBy: 1,
			Degree:       model.DegreeLikely,
		})

		assert.Equal(t, model.GamePlayer{
			ID:           2,
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
