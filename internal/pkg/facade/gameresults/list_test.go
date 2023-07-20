package gameresults

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListGameResults(t *testing.T) {
	t.Run("error. some error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResults(mock.Anything).Return(nil, errors.New("some error"))

		got, err := fx.facade.ListGameResults(fx.ctx)

		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameResultStorage.EXPECT().GetGameResults(mock.Anything).Return([]database.GameResult{
			{
				ID:       1,
				FkGameID: 1,
				Place:    1,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			},
			{
				ID:       2,
				FkGameID: 2,
				Place:    2,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			},
		}, nil)

		got, err := fx.facade.ListGameResults(fx.ctx)

		assert.Equal(t, []model.GameResult{
			{
				ID:          1,
				FkGameID:    1,
				ResultPlace: 1,
				RoundPoints: maybe.Just("{}"),
			},
			{
				ID:          2,
				FkGameID:    2,
				ResultPlace: 2,
				RoundPoints: maybe.Just("{}"),
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
