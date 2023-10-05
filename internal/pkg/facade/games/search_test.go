package games

import (
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

func TestFacade_SearchGamesByLeagueID(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: get total error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"league_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(0, errors.New("some error"))

		got, total, err := fx.facade.SearchGamesByLeagueID(fx.ctx, 1, 1, 10)
		assert.Nil(t, got)
		assert.Equal(t, uint64(0), total)
		assert.Error(t, err)
	})

	t.Run("error: find with limit error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"league_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(10, nil)

		fx.gameStorage.EXPECT().FindWithLimit(mock.Anything, builder.And(
			builder.Eq{
				"league_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "date", uint64(1), uint64(10)).Return(nil, errors.New("some error"))

		got, total, err := fx.facade.SearchGamesByLeagueID(fx.ctx, 1, 1, 10)
		assert.Nil(t, got)
		assert.Equal(t, uint64(0), total)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"league_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(10, nil)

		fx.gameStorage.EXPECT().FindWithLimit(mock.Anything, builder.And(
			builder.Eq{
				"league_id": int32(1),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "date", uint64(1), uint64(1)).Return([]database.Game{
			{
				ID:   1,
				Date: timeutils.TimeNow(),
			},
		}, nil)

		got, total, err := fx.facade.SearchGamesByLeagueID(fx.ctx, 1, 1, 1)
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow()),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
			},
		}, got)
		assert.Equal(t, uint64(10), total)
		assert.NoError(t, err)
	})
}

func TestFacade_SearchPassedAndRegisteredGames(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: total error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().Add(-3600 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(0, errors.New("some error"))

		got, total, err := fx.facade.SearchPassedAndRegisteredGames(fx.ctx, 0, 10)
		assert.Nil(t, got)
		assert.Equal(t, uint64(0), total)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().Add(-3600 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(10, nil)

		fx.gameStorage.EXPECT().FindWithLimit(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().Add(-3600 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "-date", uint64(0), uint64(1)).Return(nil, errors.New("some error"))

		got, total, err := fx.facade.SearchPassedAndRegisteredGames(fx.ctx, 0, 1)
		assert.Nil(t, got)
		assert.Equal(t, uint64(0), total)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Total(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().Add(-3600 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		)).Return(10, nil)

		fx.gameStorage.EXPECT().FindWithLimit(mock.Anything, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().Add(-3600 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "-date", uint64(0), uint64(1)).Return([]database.Game{
			{
				ID:   1,
				Date: timeutils.TimeNow().Add(-3601 * time.Second),
			},
		}, nil)

		got, total, err := fx.facade.SearchPassedAndRegisteredGames(fx.ctx, 0, 1)
		assert.Equal(t, []model.Game{
			{
				ID:          1,
				ExternalID:  maybe.Nothing[int32](),
				Name:        maybe.Nothing[string](),
				Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
				PaymentType: maybe.Nothing[string](),
				Payment:     maybe.Nothing[model.Payment](),
				HasPassed:   true,
			},
		}, got)
		assert.Equal(t, uint64(10), total)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
