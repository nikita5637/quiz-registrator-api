package gamephotos

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_AddGamePhotos(t *testing.T) {
	t.Run("error while find game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		err := fx.facade.AddGamePhotos(fx.ctx, 1, nil)
		assert.Error(t, err)
	})

	t.Run("error game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, sql.ErrNoRows)

		err := fx.facade.AddGamePhotos(fx.ctx, 1, nil)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameNotFound)
	})

	t.Run("error while insert url", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url1",
		}).Return(1, nil)
		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url2",
		}).Return(2, nil)
		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url3",
		}).Return(0, errors.New("some error"))

		err := fx.facade.AddGamePhotos(fx.ctx, 1, []string{
			"url1",
			"url2",
			"url3",
		})
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url1",
		}).Return(1, nil)
		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url2",
		}).Return(2, nil)
		fx.gamePhotoStorage.EXPECT().Insert(mock.Anything, model.GamePhoto{
			FkGameID: 1,
			URL:      "url3",
		}).Return(3, nil)

		err := fx.facade.AddGamePhotos(fx.ctx, 1, []string{
			"url1",
			"url2",
			"url3",
		})
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
