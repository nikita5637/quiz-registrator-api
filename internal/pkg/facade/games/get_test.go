package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_GetGameByID(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{}, sql.ErrNoRows)

		got, err := fx.facade.GetGameByID(fx.ctx, 1)
		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{}, errors.New("some error"))

		got, err := fx.facade.GetGameByID(fx.ctx, 1)
		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: game is deleted", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		timeNow := timeutils.TimeNow()
		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
			DeletedAt: sql.NullTime{
				Time:  timeNow,
				Valid: true,
			},
		}, nil)

		got, err := fx.facade.GetGameByID(fx.ctx, 1)
		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: game has not passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		timeNow := timeutils.TimeNow()
		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID:   1,
			Date: timeNow.Add(time.Hour),
		}, nil)

		got, err := fx.facade.GetGameByID(fx.ctx, 1)
		expectedGame := model.NewGame()
		expectedGame.ID = 1
		expectedGame.Date = model.DateTime(timeNow.Add(time.Hour))
		assert.Equal(t, expectedGame, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		timeNow := timeutils.TimeNow()
		fx.gameStorage.EXPECT().GetGameByID(mock.Anything, 1).Return(&database.Game{
			ID: 1,
			ExternalID: sql.NullInt64{
				Int64: 123,
				Valid: true,
			},
			LeagueID: 1,
			Date:     timeNow.Add(-3601 * time.Second),
		}, nil)

		got, err := fx.facade.GetGameByID(fx.ctx, 1)
		expectedGame := model.NewGame()
		expectedGame.ID = 1
		expectedGame.ExternalID = maybe.Just(int32(123))
		expectedGame.LeagueID = 1
		expectedGame.Date = model.DateTime(timeNow.Add(-3601 * time.Second))
		expectedGame.HasPassed = true
		expectedGame.GameLink = maybe.Just("https://spb.quizplease.ru/game-page?id=123")
		assert.Equal(t, expectedGame, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_GetGamesByIDs(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.In("id", []int32{1, 2, 3}),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(nil, errors.New("some error"))

		got, err := fx.facade.GetGamesByIDs(fx.ctx, []int32{1, 2, 3})
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.In("id", []int32{1, 2, 3}),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		got, err := fx.facade.GetGamesByIDs(fx.ctx, []int32{1, 2, 3})
		assert.Equal(t, []model.Game{}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.In("id", []int32{1, 2, 3}),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{
			{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 777,
					Valid: true,
				},
				LeagueID: 3,
				Date:     timeutils.TimeNow().Add(time.Hour),
			},
			{
				ID:   3,
				Date: timeutils.TimeNow().Add(-3601 * time.Second),
			},
		}, nil)

		got, err := fx.facade.GetGamesByIDs(fx.ctx, []int32{1, 2, 3})
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Just(int32(777)),
				LeagueID:    3,
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(time.Hour)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   false,
				GameLink:    maybe.Just("https://club60sec.ru/quizgames/game/777/"),
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

func TestFacade_GetTodaysGames(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2006-01-02%\""),
		), "").Return(nil, errors.New("some error"))

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Nil(t, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok: empty list", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2006-01-02%\""),
		), "").Return([]database.Game{}, nil)

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Equal(t, []model.Game{}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Expr("date LIKE \"2006-01-02%\""),
		), "").Return([]database.Game{
			{
				ID: 1,
				ExternalID: sql.NullInt64{
					Int64: 777,
					Valid: true,
				},
				LeagueID:   1,
				Date:       timeutils.TimeNow().Add(time.Hour),
				Registered: true,
			},
			{
				ID:         3,
				Date:       timeutils.TimeNow().Add(-3601 * time.Second),
				Registered: true,
			},
		}, nil)

		got, err := fx.facade.GetTodaysGames(fx.ctx)
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Just(int32(777)),
				LeagueID:    1,
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(time.Hour)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				Registered:  true,
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
				Registered:  true,
				HasPassed:   true,
				GameLink:    maybe.Nothing[string](),
			},
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
