package games

import (
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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_CreateGame(t *testing.T) {
	timeutils.TimeNow = func() time.Time {
		return timeutils.ConvertTime("2006-01-02 15:04")
	}

	t.Run("error: find already existed games error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.Eq{
				"external_id": int32(1),
				"league_id":   int32(1),
				"place_id":    int32(1),
				"number":      "1",
				"date":        timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(nil, errors.New("some error"))

		createdGame := model.NewGame()
		createdGame.ExternalID = maybe.Just(int32(1))
		createdGame.LeagueID = 1
		createdGame.PlaceID = 1
		createdGame.Number = "1"
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)
		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: game already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
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
				ID: 1,
			},
		}, nil)

		createdGame := model.NewGame()
		createdGame.LeagueID = 1
		createdGame.PlaceID = 1
		createdGame.Number = "number"
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)

		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
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
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(0),
				"place_id":  int32(0),
				"number":    "",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().CreateGame(mock.Anything, database.Game{
			Date: timeutils.TimeNow(),
		}).Return(0, &mysql.MySQLError{
			Message: leagueIBFK1ConstraintName,
			Number:  1452,
		})

		createdGame := model.NewGame()
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)

		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
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
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(0),
				"place_id":  int32(0),
				"number":    "",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().CreateGame(mock.Anything, database.Game{
			Date: timeutils.TimeNow(),
		}).Return(0, &mysql.MySQLError{
			Message: "",
			Number:  1452,
		})

		createdGame := model.NewGame()
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)

		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, places.ErrPlaceNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error: internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(0),
				"place_id":  int32(0),
				"number":    "",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().CreateGame(mock.Anything, database.Game{
			Date: timeutils.TimeNow(),
		}).Return(0, errors.New("some error"))

		createdGame := model.NewGame()
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)

		expectedGame := model.NewGame()
		assert.Equal(t, expectedGame, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.And(
			builder.IsNull{
				"external_id",
			},
			builder.Eq{
				"league_id": int32(0),
				"place_id":  int32(0),
				"number":    "",
				"date":      timeutils.TimeNow(),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "").Return([]database.Game{}, nil)

		fx.gameStorage.EXPECT().CreateGame(mock.Anything, database.Game{
			Date: timeutils.TimeNow(),
		}).Return(1, nil)

		createdGame := model.NewGame()
		createdGame.Date = model.DateTime(timeutils.TimeNow())
		got, err := fx.facade.CreateGame(fx.ctx, createdGame)

		expectedGame := model.NewGame()
		expectedGame.ID = 1
		expectedGame.Date = model.DateTime(timeutils.TimeNow())
		assert.Equal(t, expectedGame, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
