package games

import (
	"errors"
	"testing"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_AddGames(t *testing.T) {
	activeGameLag := uint16(3600)
	assert.Greater(t, activeGameLag, uint16(1))

	cfg := config.GlobalConfig{}
	cfg.ActiveGameLag = activeGameLag

	config.UpdateGlobalConfig(cfg)

	game1 := model.Game{
		ID:        1,
		LeagueID:  model.LeagueQuizPlease,
		Type:      model.GameTypeClassic,
		Number:    "1",
		PlaceID:   1,
		Date:      model.DateTime(time_utils.ConvertTime("2022-01-01 16:30")),
		DeletedAt: model.DateTime(time_utils.TimeNow()),
	}

	game2 := model.Game{
		ID:        2,
		LeagueID:  model.LeagueQuizPlease,
		Type:      model.GameTypeClassic,
		Number:    "2",
		PlaceID:   1,
		Date:      model.DateTime(time_utils.ConvertTime("2022-01-02 16:30")),
		DeletedAt: model.DateTime(time_utils.TimeNow()),
	}

	game3 := model.Game{
		ID:       3,
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeClassic,
		Number:   "1",
		PlaceID:  2,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-03 13:00")),
	}

	game4 := model.Game{
		ID:       4,
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeMoviesAndMusic,
		Number:   "1",
		PlaceID:  3,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-03 16:30")),
	}

	game5 := model.Game{
		ID:       5,
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeClassic,
		Number:   "3",
		PlaceID:  1,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-04 16:30")),
	}

	game6 := model.Game{
		ID:       6,
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeMoviesAndMusic,
		Number:   "2",
		PlaceID:  2,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-04 16:30")),
	}

	game7 := model.Game{
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeClassic,
		Number:   "4",
		PlaceID:  1,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-07 16:30")),
	}
	game8 := model.Game{
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeMoviesAndMusic,
		Number:   "2",
		PlaceID:  2,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-08 13:00")),
	}
	game9 := model.Game{
		LeagueID: model.LeagueQuizPlease,
		Type:     model.GameTypeMoviesAndMusic,
		Number:   "3",
		PlaceID:  3,
		Date:     model.DateTime(time_utils.ConvertTime("2022-01-08 16:30")),
	}

	storageGames := []model.Game{
		game1,
		game2,
		game3,
		game4,
		game5,
		game6,
	}

	t.Run("test case 1", func(t *testing.T) {
		// both deleted games not active
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:35")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game5,
			game6,
			game7,
			game8,
			game9,
		}
		newGames[0].ID = 0
		newGames[1].ID = 0

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{
			storageGames[4],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{
			storageGames[5],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[3]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[3]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[4]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[4]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(4)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 2", func(t *testing.T) {
		// first game is inactive
		// no second game
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 16:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game5,
			game6,
			game7,
			game8,
			game9,
		}
		newGames[0].ID = 0
		newGames[1].ID = 0

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{
			storageGames[4],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{
			storageGames[5],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[3]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[3]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[4]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[4]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(4)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 3", func(t *testing.T) {
		// first and second games is active
		// no third game
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 12:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game3,
			game4,
			game5,
			game7,
			game8,
			game9,
		}
		newGames[0].ID = 0

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{
			storageGames[2],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{
			storageGames[3],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{
			storageGames[4],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[3]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[3]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[4]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[4]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[5]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[5]).Return(8, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(6)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 4", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:29")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game5,
			game6,
			game7,
			game8,
			game9,
		}
		newGames[0].ID = 0
		newGames[1].ID = 0

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{
			storageGames[4],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{
			storageGames[5],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[3]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[3]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[4]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[4]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 5", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:30")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game5,
			game6,
			game7,
			game8,
			game9,
		}
		newGames[0].ID = 0
		newGames[1].ID = 0

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{
			storageGames[4],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{
			storageGames[5],
		}, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[3]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[3]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[4]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[4]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(4)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 6", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 12:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[1]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(4)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(5)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(6)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("test case 7", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[1]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(5)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(6)).Return(nil)

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("get not deleted games error ", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(nil, errors.New("some error"))

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while get game", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return(nil, errors.New("some error"))

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while insert game", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[1]).Return(8, errors.New("some error"))

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while delete game that not contain in master", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[1]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(5)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(6)).Return(errors.New("some error"))

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error while delete outdated game", func(t *testing.T) {
		fx := tearUp(t)
		time_utils.TimeNow = func() time.Time {
			return time_utils.ConvertTime("2022-01-03 17:00")
		}

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		notDeletedGames := storageGames[2:]

		newGames := []model.Game{
			game7,
			game8,
			game9,
		}

		fx.gameStorage.EXPECT().Find(mock.Anything, builder.NewCond().And(
			builder.In(
				"league_id", []int32{
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
					model.LeagueQuizPlease,
				},
			),
			builder.IsNull{
				"deleted_at",
			},
		), "").Return(notDeletedGames, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[0]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[0]).Return(7, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[1]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[1]).Return(8, nil)

		fx.gameStorage.EXPECT().Find(mock.Anything, getGamesStorageFindBuilder(newGames[2]), "").Return([]model.Game{}, nil)
		fx.gameStorage.EXPECT().Insert(mock.Anything, newGames[2]).Return(9, nil)

		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(5)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(6)).Return(nil)
		fx.gameStorage.EXPECT().Delete(mock.Anything, int32(3)).Return(errors.New("some error"))

		err := fx.facade.AddGames(fx.ctx, newGames)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func getGamesStorageFindBuilder(game model.Game) builder.Cond {
	return builder.NewCond().And(
		builder.Eq{
			"league_id": game.LeagueID,
			"type":      game.Type,
			"number":    game.Number,
			"place_id":  game.PlaceID,
			"date":      game.DateTime().AsTime(),
		},
		builder.IsNull{
			"deleted_at",
		},
	)
}

func Test_getGameIDsForDelete(t *testing.T) {
	time_utils.TimeNow = func() time.Time {
		return time_utils.ConvertTime("2022-01-03 15:00")
	}

	type args struct {
		games       []model.Game
		activeGames []model.Game
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "tc1",
			args: args{
				games: []model.Game{
					{
						ID: 1,
					},
					{
						ID: 2,
					},
					{
						ID: 3,
					},
				},
				activeGames: []model.Game{
					{
						ID:   3,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 13:00")),
					},
					{
						ID:   4,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 15:00")),
					},
					{
						ID:   5,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 16:30")),
					},
				},
			},
			want: []int32{
				4,
				5,
			},
		},
		{
			name: "tc2",
			args: args{
				games: []model.Game{
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
				},
				activeGames: []model.Game{},
			},
			want: []int32{},
		},
		{
			name: "tc3",
			args: args{
				games: []model.Game{},
				activeGames: []model.Game{
					{
						ID:   3,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 16:00")),
					},
					{
						ID:   4,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 17:00")),
					},
					{
						ID:   5,
						Date: model.DateTime(time_utils.ConvertTime("2022-01-03 16:30")),
					},
				},
			},
			want: []int32{
				3,
				4,
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getGameIDsForDelete(tt.args.games, tt.args.activeGames)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
