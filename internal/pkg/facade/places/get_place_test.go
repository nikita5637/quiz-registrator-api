package places

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetPlaceByID(t *testing.T) {
	t.Run("sql.ErrNoRows error", func(t *testing.T) {
		fx := tearUp(t)

		fx.placeStorage.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{}, sql.ErrNoRows)

		got, err := fx.facade.GetPlaceByID(fx.ctx, 1)
		assert.Equal(t, model.Place{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPlaceNotFound)
	})

	t.Run("internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.placeStorage.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{}, errors.New("some error"))

		got, err := fx.facade.GetPlaceByID(fx.ctx, 1)
		assert.Equal(t, model.Place{}, got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.placeStorage.EXPECT().GetPlaceByID(fx.ctx, int32(1)).Return(model.Place{
			ID:        1,
			Name:      "place name",
			Address:   "place address",
			ShortName: "short name",
			Latitude:  1.1,
			Longitude: 2.2,
			MenuLink:  "menu link",
		}, nil)

		got, err := fx.facade.GetPlaceByID(fx.ctx, 1)
		assert.Equal(t, model.Place{
			ID:        1,
			Name:      "place name",
			Address:   "place address",
			ShortName: "short name",
			Latitude:  1.1,
			Longitude: 2.2,
			MenuLink:  "menu link",
		}, got)
		assert.NoError(t, err)
	})
}
