package games

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/builder"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchGame(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	viper.Set("service.game.has_passed_game_lag", 3600)

	t.Run("error: find error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"external_id": int32(1),
				"league_id":   int32(1),
				"place_id":    int32(1),
				"number":      "number",
				"date":        timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(nil, errors.New("some error"))

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:         1,
			ExternalID: maybe.Just(int32(1)),
			LeagueID:   1,
			PlaceID:    1,
			Number:     "number",
			Date:       model.DateTime(timeutils.TimeNow()),
		})
		assert.Equal(t, model.NewGame(), got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: game already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(1),
				"place_id":  int32(1),
				"number":    "number",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{
			{
				ID:       1,
				LeagueID: 1,
			},
		}, nil)

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow()),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.NewGame(), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrGameAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: league not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(1),
				"place_id":  int32(1),
				"number":    "number",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().PatchGame(mock.Anything, database.Game{
			ID:       1,
			LeagueID: 1,
			PlaceID:  1,
			Number:   "number",
			Date:     timeutils.TimeNow(),
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: leagueIBFK1ConstraintName,
		})

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow()),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.NewGame(), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, leagues.ErrLeagueNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: place not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(1),
				"place_id":  int32(1),
				"number":    "number",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().PatchGame(mock.Anything, database.Game{
			ID:       1,
			LeagueID: 1,
			PlaceID:  1,
			Number:   "number",
			Date:     timeutils.TimeNow(),
		}).Return(&mysql.MySQLError{
			Number:  1452,
			Message: "",
		})

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow()),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.NewGame(), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, places.ErrPlaceNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: patch error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(1),
				"place_id":  int32(1),
				"number":    "number",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().PatchGame(mock.Anything, database.Game{
			ID:       1,
			LeagueID: 1,
			PlaceID:  1,
			Number:   "number",
			Date:     timeutils.TimeNow(),
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow()),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.NewGame(), got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok with externalID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.Eq{
				"external_id": int32(777),
				"league_id":   int32(1),
				"place_id":    int32(1),
				"number":      "number",
				"date":        timeutils.TimeNow().Add(-3601 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().PatchGame(mock.Anything, database.Game{
			ID: 1,
			ExternalID: sql.NullInt64{
				Int64: 777,
				Valid: true,
			},
			LeagueID: 1,
			PlaceID:  1,
			Number:   "number",
			Date:     timeutils.TimeNow().Add(-3601 * time.Second),
		}).Return(nil)

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Just(int32(777)),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.Game{
			ID:          1,
			ExternalID:  maybe.Just(int32(777)),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
			HasPassed:   true,
			GameLink:    maybe.Just("https://spb.quizplease.ru/game-page?id=777"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok without externalID", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Neq{
				"id": int32(1),
			},
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(1),
				"place_id":  int32(1),
				"number":    "number",
				"date":      timeutils.TimeNow().Add(-3601 * time.Second),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().PatchGame(mock.Anything, database.Game{
			ID:       1,
			LeagueID: 1,
			PlaceID:  1,
			Number:   "number",
			Date:     timeutils.TimeNow().Add(-3601 * time.Second),
		}).Return(nil)

		got, err := fx.facade.PatchGame(fx.ctx, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
		})
		assert.Equal(t, model.Game{
			ID:          1,
			ExternalID:  maybe.Nothing[int32](),
			LeagueID:    1,
			Number:      "number",
			Name:        maybe.Nothing[string](),
			PlaceID:     1,
			Date:        model.DateTime(timeutils.TimeNow().Add(-3601 * time.Second)),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[model.Payment](),
			HasPassed:   true,
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
