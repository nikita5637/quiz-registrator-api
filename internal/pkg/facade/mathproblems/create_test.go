package mathproblems

import (
	"errors"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateMathProblem(t *testing.T) {
	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.mathProblemStorage.EXPECT().CreateMathProblem(mock.Anything, database.MathProblem{
			FkGameID: 1,
			URL:      "url",
		}).Return(0, &mysql.MySQLError{
			Number:  1452,
			Message: mathProblemIBFK1,
		})

		got, err := fx.facade.CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "url",
		})
		assert.Equal(t, model.MathProblem{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: math problem already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.mathProblemStorage.EXPECT().CreateMathProblem(mock.Anything, database.MathProblem{
			FkGameID: 1,
			URL:      "url",
		}).Return(0, &mysql.MySQLError{
			Number: 1062,
		})

		got, err := fx.facade.CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "url",
		})
		assert.Equal(t, model.MathProblem{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrMathProblemAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.mathProblemStorage.EXPECT().CreateMathProblem(mock.Anything, database.MathProblem{
			FkGameID: 1,
			URL:      "url",
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "url",
		})
		assert.Equal(t, model.MathProblem{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.mathProblemStorage.EXPECT().CreateMathProblem(mock.Anything, database.MathProblem{
			FkGameID: 1,
			URL:      "url",
		}).Return(1, nil)

		got, err := fx.facade.CreateMathProblem(fx.ctx, model.MathProblem{
			GameID: 1,
			URL:    "url",
		})
		assert.Equal(t, model.MathProblem{
			ID:     1,
			GameID: 1,
			URL:    "url",
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
