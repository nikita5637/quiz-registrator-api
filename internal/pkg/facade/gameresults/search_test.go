package gameresults

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_SearchGameResultByGameID(t *testing.T) {
	t.Run("error: sql.ErrNoRows", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(mock.Anything, 1).Return(model.GameResult{}, sql.ErrNoRows)

		got, err := fx.facade.SearchGameResultByGameID(fx.ctx, int32(1))
		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameResultNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(mock.Anything, 1).Return(model.GameResult{}, errors.New("some error"))

		got, err := fx.facade.SearchGameResultByGameID(fx.ctx, int32(1))
		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(mock.Anything, 1).Return(model.GameResult{
			ID:          1,
			RoundPoints: maybe.Nothing[string](),
		}, nil)

		got, err := fx.facade.SearchGameResultByGameID(fx.ctx, int32(1))
		assert.Equal(t, model.GameResult{
			ID:          1,
			RoundPoints: maybe.Nothing[string](),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
