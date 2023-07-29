package gameresults

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFacade_PatchGameResult(t *testing.T) {
	t.Run("get original game result error. game result not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, sql.ErrNoRows)

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameResultNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("get original game result error. some error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, errors.New("some error"))

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("patch game result error. game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, nil)

		fx.gameResultStorage.EXPECT().PatchGameResult(mock.Anything, database.GameResult{
			ID:       1,
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{\"round1\":6}",
			},
		}).Return(&mysql.MySQLError{
			Number: 1452,
		})

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, games.ErrGameNotFound)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("patch game result error. game result already exists", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, nil)

		fx.gameResultStorage.EXPECT().PatchGameResult(mock.Anything, database.GameResult{
			ID:       1,
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{\"round1\":6}",
			},
		}).Return(&mysql.MySQLError{
			Number: 1062,
		})

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrGameResultAlreadyExists)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("patch game result error. some MySQL error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, nil)

		fx.gameResultStorage.EXPECT().PatchGameResult(mock.Anything, database.GameResult{
			ID:       1,
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{\"round1\":6}",
			},
		}).Return(&mysql.MySQLError{
			Number: 1,
		})

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("patch game result error. some error", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectRollback()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, nil)

		fx.gameResultStorage.EXPECT().PatchGameResult(mock.Anything, database.GameResult{
			ID:       1,
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{\"round1\":6}",
			},
		}).Return(errors.New("some error"))

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{}, got)
		assert.Error(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.dbMock.ExpectBegin()
		fx.dbMock.ExpectCommit()

		fx.gameResultStorage.EXPECT().GetGameResultByID(mock.Anything, 1).Return(
			&database.GameResult{
				ID:       1,
				FkGameID: 1,
				Place:    3,
				Points: sql.NullString{
					Valid:  true,
					String: "{}",
				},
			}, nil)

		fx.gameResultStorage.EXPECT().PatchGameResult(mock.Anything, database.GameResult{
			ID:       1,
			FkGameID: 2,
			Place:    1,
			Points: sql.NullString{
				Valid:  true,
				String: "{\"round1\":6}",
			},
		}).Return(nil)

		got, err := fx.facade.PatchGameResult(fx.ctx, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, []string{
			fieldNameGameID,
			fieldNameResultPlace,
			fieldNameRoundPoints,
		})

		assert.Equal(t, model.GameResult{
			ID:          1,
			FkGameID:    2,
			ResultPlace: 1,
			RoundPoints: maybe.Just("{\"round1\":6}"),
		}, got)
		assert.NoError(t, err)

		err = fx.dbMock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}

func TestFacade_checkPathNames(t *testing.T) {
	field, _ := reflect.ValueOf(gameresultmanagerpb.GameResult{}).Type().FieldByName("GameId")
	assert.Equal(t, fieldNameGameID, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(gameresultmanagerpb.GameResult{}).Type().FieldByName("ResultPlace")
	assert.Equal(t, fieldNameResultPlace, strings.Split(field.Tag.Get("json"), ",")[0])
	field, _ = reflect.ValueOf(gameresultmanagerpb.GameResult{}).Type().FieldByName("RoundPoints")
	assert.Equal(t, fieldNameRoundPoints, strings.Split(field.Tag.Get("json"), ",")[0])
	assert.Equal(t, 7, reflect.ValueOf(gameresultmanagerpb.GameResult{}).Type().NumField())
}
