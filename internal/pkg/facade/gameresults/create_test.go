package gameresults

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateGameResult(t *testing.T) {
	t.Run("create game result error. game result already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().CreateGameResult(mock.Anything, database.GameResult{
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{}",
			},
		}).Return(0, &mysql.MySQLError{
			Number: 1062,
		})

		got, err := fx.facade.CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameResultAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("create game result error. other error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().CreateGameResult(mock.Anything, database.GameResult{
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{}",
			},
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameResultStorage.EXPECT().CreateGameResult(mock.Anything, database.GameResult{
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{}",
			},
		}).Return(1, nil)

		got, err := fx.facade.CreateGameResult(fx.ctx, model.GameResult{
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: model.NewMaybeString("{}"),
		})

		assert.Equal(t, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: model.NewMaybeString("{}"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
