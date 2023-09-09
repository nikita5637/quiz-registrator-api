package gamephotos

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFacade_GetGameWithPhotosIDs(t *testing.T) {
	t.Run("some error while getting game IDs", func(t *testing.T) {
		fx := tearUp(t)
		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 0, 0)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("empty game IDs list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return(nil, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 0, 0)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("offset gt len(GetGameIDsWithPhotos)", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 0, 6)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("offset ge then len(GetGameIDsWithPhotos)", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 0, 5)
		assert.NoError(t, err)
		assert.Nil(t, got)
	})

	t.Run("get middle batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 2, 2)
		assert.NoError(t, err)

		assert.ElementsMatch(t, got, []int32{
			3,
			4,
		})
	})

	t.Run("get middle batch #2 without limit", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 0, 2)
		assert.NoError(t, err)
		assert.ElementsMatch(t, got, []model.Game{})
	})

	t.Run("get last batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 2, 3)
		assert.NoError(t, err)

		assert.ElementsMatch(t, got, []int32{4, 5})
	})

	t.Run("get last batch #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 5, 3)
		assert.NoError(t, err)

		assert.ElementsMatch(t, got, []int32{4, 5})
	})

	t.Run("get first batch #1", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 3, 0)
		assert.NoError(t, err)

		assert.ElementsMatch(t, got, []int32{1, 2, 3})
	})

	t.Run("get first batch #2", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetGameWithPhotosIDs(fx.ctx, 30, 0)
		assert.NoError(t, err)

		assert.ElementsMatch(t, got, []int32{1, 2, 3, 4, 5})
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

		fx.gamePhotoStorage.EXPECT().GetGameIDsWithPhotos(fx.ctx, uint32(0)).Return([]int{1, 2, 3, 4, 5}, nil)

		got, err := fx.facade.GetNumberOfGamesWithPhotos(fx.ctx)
		assert.NoError(t, err)
		assert.Equal(t, uint32(5), got)
	})
}

func TestFacade_GetPhotosByGameID(t *testing.T) {
	t.Run("error while find game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{}, errors.New("some error"))

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("error game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{}, sql.ErrNoRows)

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)
	})

	t.Run("error while get game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("empty url list", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, 1).Return([]*database.GamePhoto{}, nil)

		got, err := fx.facade.GetPhotosByGameID(fx.ctx, 1)
		assert.Equal(t, []string{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(fx.ctx, 1).Return([]*database.GamePhoto{
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
