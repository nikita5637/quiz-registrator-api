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
	"github.com/stretchr/testify/mock"
)

func TestFacade_IsGameHasPhotos(t *testing.T) {
	t.Run("error while find game", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{}, errors.New("some error"))

		got, err := fx.facade.IsGameHasPhotos(fx.ctx, 1)
		assert.False(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{}, sql.ErrNoRows)

		got, err := fx.facade.IsGameHasPhotos(fx.ctx, 1)
		assert.False(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while get game photos", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(mock.Anything, 1).Return(nil, errors.New("some error"))

		got, err := fx.facade.IsGameHasPhotos(fx.ctx, 1)
		assert.False(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: write logs error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(mock.Anything, 1).Return([]*database.GamePhoto{}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotIndicationIfGameHasPhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(errors.New("some error"))

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		got, err := fx.facade.IsGameHasPhotos(ctx, 1)
		assert.False(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok false", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(mock.Anything, 1).Return([]*database.GamePhoto{}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotIndicationIfGameHasPhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(nil)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		got, err := fx.facade.IsGameHasPhotos(ctx, 1)
		assert.False(t, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok true", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
		}, nil)

		fx.gamePhotoStorage.EXPECT().GetGamePhotosByGameID(mock.Anything, 1).Return([]*database.GamePhoto{
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

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotIndicationIfGameHasPhotos,
			ObjectType: maybe.Just(quizlogger.ObjectTypeGame),
			ObjectID:   maybe.Just(int32(1)),
			Metadata:   nil,
		}).Return(nil)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		got, err := fx.facade.IsGameHasPhotos(ctx, 1)
		assert.True(t, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
