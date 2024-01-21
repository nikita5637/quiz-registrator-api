package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_ListGames(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return(nil, errors.New("some error"))

		got, err := fx.facade.ListGames(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: write logs error", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return([]database.Game{}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCompleteListOfGames,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}).Return(errors.New("some error"))

		got, err := fx.facade.ListGames(ctx)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return([]database.Game{}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCompleteListOfGames,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}).Return(nil)

		got, err := fx.facade.ListGames(ctx)
		assert.Equal(t, []model.Game{}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		ctx := usersutils.NewContextWithUser(fx.ctx, &model.User{
			ID: 1,
		})

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date").Return([]database.Game{
			{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 777,
					Valid: true,
				},
				LeagueID: 1,
				Date:     timeutils.TimeNow().Add(time.Hour),
			},
			{
				ID:   3,
				Date: timeutils.TimeNow().Add(-3601 * time.Second),
			},
		}, nil)

		fx.quizLogger.EXPECT().Write(mock.Anything, quizlogger.Params{
			UserID:     maybe.Just(int32(1)),
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCompleteListOfGames,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}).Return(nil)

		got, err := fx.facade.ListGames(ctx)
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Just(int32(777)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(time.Hour)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   false,
				GameLink:    maybe.Just("https://spb.quizplease.ru/game-page?id=777"),
			},
			{
				ID:          3,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   true,
				GameLink:    maybe.Nothing[string](),
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
