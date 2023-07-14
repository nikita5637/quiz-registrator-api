package gamephotos

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetGamesWithPhotos(t *testing.T) {
	t.Run("some error while getting game IDs", func(t *testing.T) {
		fx := tearUp(t)
		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 0, 0)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("empty game IDs list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return(nil, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 0, 0)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("offset g then len(GetGameIDsWithPhotos)", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 0, 6)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("offset ge then len(GetGameIDsWithPhotos)", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 0, 5)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("get middle batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 1,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    3,
			ResultPlace: 2,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 2, 2)
		assert.NoError(t, err)

		expected := []model.Game{
			{
				ID: 3,
			},
			{
				ID: 4,
			},
		}
		expected[0].ResultPlace = 1
		expected[1].ResultPlace = 2

		assert.ElementsMatch(t, got, expected)
	})

	t.Run("get middle batch #2 without limit", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 0, 2)
		assert.NoError(t, err)
		assert.ElementsMatch(t, got, []model.Game{})
	})

	t.Run("get middle batch #3 with error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, errors.New("some error"))

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 1,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 2, 2)
		assert.Error(t, err)
		assert.ElementsMatch(t, got, nil)
	})

	t.Run("get last batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(5)).Return(model.Game{
			ID: 5,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    4,
			ResultPlace: 4,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 5).Return(model.GameResult{
			ID:          5,
			FkGameID:    5,
			ResultPlace: 5,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 2, 3)
		assert.NoError(t, err)

		expected := []model.Game{
			{
				ID: 4,
			},
			{
				ID: 5,
			},
		}
		expected[0].ResultPlace = 4
		expected[1].ResultPlace = 5

		assert.ElementsMatch(t, got, expected)
	})

	t.Run("get last batch #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(5)).Return(model.Game{
			ID: 5,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    4,
			ResultPlace: 4,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 5).Return(model.GameResult{
			ID:          5,
			FkGameID:    5,
			ResultPlace: 5,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 5, 3)
		assert.NoError(t, err)

		expected := []model.Game{
			{
				ID: 4,
			},
			{
				ID: 5,
			},
		}
		expected[0].ResultPlace = 4
		expected[1].ResultPlace = 5

		assert.ElementsMatch(t, got, expected)
	})

	t.Run("get last batch #3 with error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(5)).Return(model.Game{}, errors.New("some error"))

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    4,
			ResultPlace: 4,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 5, 3)
		assert.Error(t, err)
		assert.ElementsMatch(t, got, nil)
	})

	t.Run("get first batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			ID: 2,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 1).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 2).Return(model.GameResult{
			ID:          2,
			FkGameID:    2,
			ResultPlace: 2,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 3,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 3, 0)
		assert.NoError(t, err)

		expected := []model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
			{
				ID: 3,
			},
		}
		expected[0].ResultPlace = 1
		expected[1].ResultPlace = 2
		expected[2].ResultPlace = 3

		assert.ElementsMatch(t, got, expected)
	})

	t.Run("get first batch #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			ID: 2,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(5)).Return(model.Game{
			ID: 5,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 1).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 2).Return(model.GameResult{
			ID:          2,
			FkGameID:    2,
			ResultPlace: 2,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 3,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    4,
			ResultPlace: 4,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 5).Return(model.GameResult{
			ID:          5,
			FkGameID:    5,
			ResultPlace: 5,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 30, 0)
		assert.NoError(t, err)

		expected := []model.Game{
			{
				ID: 1,
			},
			{
				ID: 2,
			},
			{
				ID: 3,
			},
			{
				ID: 4,
			},
			{
				ID: 5,
			},
		}
		expected[0].ResultPlace = 1
		expected[1].ResultPlace = 2
		expected[2].ResultPlace = 3
		expected[3].ResultPlace = 4
		expected[4].ResultPlace = 5

		assert.ElementsMatch(t, got, expected)
	})

	t.Run("get first batch #3 with error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			ID: 2,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{}, errors.New("some error"))

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 1).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 2).Return(model.GameResult{
			ID:          2,
			FkGameID:    2,
			ResultPlace: 2,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 3,
		}, nil)

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 30, 0)
		assert.Error(t, err)
		assert.ElementsMatch(t, got, nil)
	})

	t.Run("get first batch with error while get game result", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(2)).Return(model.Game{
			ID: 2,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(3)).Return(model.Game{
			ID: 3,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(4)).Return(model.Game{
			ID: 4,
		}, nil)
		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(5)).Return(model.Game{
			ID: 5,
		}, nil)

		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 1).Return(model.GameResult{
			ID:          1,
			FkGameID:    1,
			ResultPlace: 1,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 2).Return(model.GameResult{
			ID:          2,
			FkGameID:    2,
			ResultPlace: 2,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 3).Return(model.GameResult{
			ID:          3,
			FkGameID:    3,
			ResultPlace: 3,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 4).Return(model.GameResult{
			ID:          4,
			FkGameID:    4,
			ResultPlace: 4,
		}, nil)
		fx.gameResultStorage.EXPECT().GetGameResultByFkGameID(fx.ctx, 5).Return(model.GameResult{}, errors.New("some error"))

		got, err := fx.facade.GetGamesWithPhotos(fx.ctx, 30, 0)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}

func TestFacade_GetNumberOfGamesWithPhotos(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetNumberOfGamesWithPhotos(fx.ctx)
		assert.Error(t, err)
		assert.Equal(t, uint32(0), got)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int32{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetNumberOfGamesWithPhotos(fx.ctx)
		assert.NoError(t, err)
		assert.Equal(t, uint32(5), got)
	})
}

func TestFacade_GetPhotosByGameID(t *testing.T) {
	t.Run("error while find game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("error while get game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, int32(1)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("empty url list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, int32(1)).Return([]model.GamePhoto{}, nil)

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Equal(t, []string{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, int32(1)).Return([]model.GamePhoto{
			{
				FkGameID: 1,
				URL:      "url1",
			},
			{
				FkGameID: 1,
				URL:      "url2",
			},
			{
				FkGameID: 1,
				URL:      "url3",
			},
		}, nil)

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.ElementsMatch(t, []string{
			"url1",
			"url2",
			"url3",
		}, got)
		assert.NoError(t, err)
	})
}
