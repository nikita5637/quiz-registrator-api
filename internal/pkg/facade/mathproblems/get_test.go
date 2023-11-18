package mathproblems

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetMathProblemByGameID(t *testing.T) {
	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.mathProblemStorage.EXPECT().GetMathProblemByGameID(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetMathProblemByGameID(fx.ctx, 1)
		assert.Equal(t, model.MathProblem{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: math prolem not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.mathProblemStorage.EXPECT().GetMathProblemByGameID(mock.Anything, 1).Return([]*database.MathProblem{}, nil)

		got, err := fx.facade.GetMathProblemByGameID(fx.ctx, 1)
		assert.Equal(t, model.MathProblem{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrMathProblemNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.mathProblemStorage.EXPECT().GetMathProblemByGameID(mock.Anything, 1).Return([]*database.MathProblem{
			{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		}, nil)

		got, err := fx.facade.GetMathProblemByGameID(fx.ctx, 1)
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
