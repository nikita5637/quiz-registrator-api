package games

import (
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetGameByID(t *testing.T) {
	// TODO
}

func TestFacade_GetGames(t *testing.T) {
	// TODO tests
}

func TestFacade_GetGamesByUserID(t *testing.T) {
	// TODO tests
}

func TestFacade_GetRegisteredGames(t *testing.T) {
	// TODO tests
}

func TestFacade_GetTodaysGames(t *testing.T) {
	t.Run("internal error while find games", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}

		fx := tearUp(t)
		fx.gameStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2022-02-10%\""),
		), "").Return(nil, errors.New("some error"))

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)
	})
	t.Run("ok", func(t *testing.T) {
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-02-10 15:31")
		}

		fx := tearUp(t)
		fx.gameStorage.EXPECT().Find(fx.ctx, builder.NewCond().And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2022-02-10%\""),
		), "").Return([]model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
		}, nil)

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Len(t, got, 2)
		assert.NoError(t, err)
	})
}
