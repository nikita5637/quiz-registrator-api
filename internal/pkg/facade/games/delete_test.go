package games

import (
	"database/sql"
	"errors"
	"testing"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_DeleteGame(t *testing.T) {
	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(nil, sql.ErrNoRows)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: get game error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(nil, errors.New("some error"))

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: game is deleted", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
			DeletedAt: sql.NullTime{
				Time:  timeutils.TimeNow(),
				Valid: true,
			},
		}, nil)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: delete game", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().DeleteGame(mock.Anything, 1).Return(errors.New("some error"))

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().DeleteGame(mock.Anything, 1).Return(nil)

		err := fx.facade.DeleteGame(fx.ctx, 1)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
