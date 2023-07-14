package leagues

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetLeagueByID(t *testing.T) {
	t.Run("sql.ErrNoRows error", func(t *testing.T) {
		fx := tearUp(t)

		fx.leagueStorage.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{}, sql.ErrNoRows)

		got, err := fx.facade.GetLeagueByID(fx.ctx, 1)
		assert.Equal(t, model.League{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrLeagueNotFound)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.leagueStorage.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{}, errors.New("some error"))

		got, err := fx.facade.GetLeagueByID(fx.ctx, 1)
		assert.Equal(t, model.League{}, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.leagueStorage.EXPECT().GetLeagueByID(fx.ctx, int32(1)).Return(model.League{
			ID:        1,
			Name:      "league name",
			ShortName: "short name",
			LogoLink:  "logo link",
			WebSite:   "web site",
		}, nil)

		got, err := fx.facade.GetLeagueByID(fx.ctx, 1)
		assert.Equal(t, model.League{
			ID:        1,
			Name:      "league name",
			ShortName: "short name",
			LogoLink:  "logo link",
			WebSite:   "web site",
		}, got)
		assert.NoError(t, err)
	})
}
