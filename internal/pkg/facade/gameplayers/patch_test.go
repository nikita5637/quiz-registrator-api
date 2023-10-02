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

func TestFacade_PatchGamePlayer(t *testing.T) {
	t.Run("Find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(nil, errors.New("some error"))

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("game player already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{
			{
				ID:       2,
				FkGameID: 2,
				FkUserID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
			},
		}, nil)

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user id not found error while patching game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		fx.gamePlayerStorage.EXPECT().PatchGamePlayer(mock.Anything, database.GamePlayer{
			ID:       1,
			FkGameID: 2,
			FkUserID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			RegisteredBy: 2,
			Degree:       2,
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: gamePlayerIBFK1ConstraintName,
		})

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, users.ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("user id for field \"registered_by\" not found error while patching game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		fx.gamePlayerStorage.EXPECT().PatchGamePlayer(mock.Anything, database.GamePlayer{
			ID:       1,
			FkGameID: 2,
			FkUserID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			RegisteredBy: 3,
			Degree:       2,
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: gamePlayerIBFK3ConstraintName,
		})

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 3,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, users.ErrUserNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("game id not found error while patching game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		fx.gamePlayerStorage.EXPECT().PatchGamePlayer(mock.Anything, database.GamePlayer{
			ID:       1,
			FkGameID: 2,
			FkUserID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			RegisteredBy: 2,
			Degree:       2,
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: gamePlayerIBFK2ConstraintName,
		})

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("internal error while patching game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		fx.gamePlayerStorage.EXPECT().PatchGamePlayer(mock.Anything, database.GamePlayer{
			ID:       1,
			FkGameID: 2,
			FkUserID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			RegisteredBy: 2,
			Degree:       2,
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"fk_game_id": int32(2),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		fx.gamePlayerStorage.EXPECT().PatchGamePlayer(mock.Anything, database.GamePlayer{
			ID:       1,
			FkGameID: 2,
			FkUserID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			RegisteredBy: 2,
			Degree:       2,
		}).Return(nil)

		got, err := fx.facade.PatchGamePlayer(fx.ctx, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		})
		assert.Equal(t, model.GamePlayer{
			ID:           1,
			GameID:       2,
			UserID:       maybe.Just(int32(2)),
			RegisteredBy: 2,
			Degree:       model.DegreeUnlikely,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
