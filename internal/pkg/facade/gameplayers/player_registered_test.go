package gameplayers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-xorm/builder"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PlayerRegisteredOnGame(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"fk_game_id": int32(1),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(nil, errors.New("some error"))

		got, err := fx.facade.PlayerRegisteredOnGame(fx.ctx, 1, 2)
		assert.False(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok. false", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"fk_game_id": int32(1),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{}, nil)

		got, err := fx.facade.PlayerRegisteredOnGame(fx.ctx, 1, 2)
		assert.False(t, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok. true", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"fk_game_id": int32(1),
				"fk_user_id": int32(2),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return([]database.GamePlayer{
			{
				ID:       1,
				FkGameID: 1,
				FkUserID: sql.NullInt64{
					Int64: 2,
					Valid: true,
				},
			},
		}, nil)

		got, err := fx.facade.PlayerRegisteredOnGame(fx.ctx, 1, 2)
		assert.True(t, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
