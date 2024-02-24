package gamephotos

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/stretchr/testify/assert"
)

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

	t.Run("error: write logs error", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.gameStorage.EXPECT().GetGameByID(ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(ctx, 1).Return([]*database.GamePhoto{}, nil)

		fx.quizLogger.EXPECT().Write(ctx, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotGamePhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(errors.New("some error"))

		got, err := fx.facade.GetPhotosByGameID(ctx, 1)
		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("empty url list", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.gameStorage.EXPECT().GetGameByID(ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(ctx, 1).Return([]*database.GamePhoto{}, nil)

		fx.quizLogger.EXPECT().Write(ctx, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotGamePhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(nil)

		got, err := fx.facade.GetPhotosByGameID(ctx, 1)
		assert.Equal(t, []string{}, got)
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.gameStorage.EXPECT().GetGameByID(ctx, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(ctx, 1).Return([]*database.GamePhoto{
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

		fx.quizLogger.EXPECT().Write(ctx, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotGamePhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(nil)

		got, err := fx.facade.GetPhotosByGameID(ctx, 1)
		assert.ElementsMatch(t, []string{
			"url1",
			"url2",
			"url3",
		}, got)
		assert.NoError(t, err)
	})
}
