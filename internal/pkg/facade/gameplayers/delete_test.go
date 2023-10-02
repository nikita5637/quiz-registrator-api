package gameplayers

import (
	"database/sql"
	"errors"
	"testing"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_DeleteGamePlayer(t *testing.T) {
	t.Run("error. game player not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{}, sql.ErrNoRows)

		err := fx.facade.DeleteGamePlayer(fx.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("some error while getting game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{}, errors.New("some error"))

		err := fx.facade.DeleteGamePlayer(fx.ctx, 1)

		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("try to delete deleted game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{
			ID: 1,
			DeletedAt: sql.NullTime{
				Valid: true,
				Time:  timeutils.TimeNow(),
			},
		}, nil)

		err := fx.facade.DeleteGamePlayer(fx.ctx, 1)

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGamePlayerNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while deleting game player", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{
			ID: 1,
		}, nil)

		fx.gamePlayerStorage.EXPECT().DeleteGamePlayer(mock.Anything, 1).Return(errors.New("some error"))

		err := fx.facade.DeleteGamePlayer(fx.ctx, 1)

		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gamePlayerStorage.EXPECT().GetGamePlayer(mock.Anything, 1).Return(&database.GamePlayer{
			ID: 1,
		}, nil)

		fx.gamePlayerStorage.EXPECT().DeleteGamePlayer(mock.Anything, 1).Return(nil)

		err := fx.facade.DeleteGamePlayer(fx.ctx, 1)

		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
