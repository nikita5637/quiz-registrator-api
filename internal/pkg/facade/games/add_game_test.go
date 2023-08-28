package games

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"
)

func TestFacade_AddGame(t *testing.T) {
	t.Run("league id validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: -1,
			Number:   "1",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidLeagueID)
	})

	t.Run("game type validation error case 1", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     0,
			Number:   "1",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidGameType)
	})

	t.Run("game type validation error case 2", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     3,
			Number:   "1",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidGameType)
	})

	t.Run("game number validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     model.GameTypeClassic,
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidGameNumber)
	})

	t.Run("place id validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     model.GameTypeClassic,
			Number:   "1",
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidPlaceID)
	})

	t.Run("date validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     model.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidDate)
	})

	t.Run("price validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     model.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
			Date:     model.DateTime(time_utils.TimeNow()),
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidPrice)
	})

	t.Run("max players validation error", func(t *testing.T) {
		fx := tearUp(t)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:     model.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
			Date:     model.DateTime(time_utils.TimeNow()),
			Price:    400,
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, model.ErrInvalidMaxPlayers)
	})

	t.Run("insert error", func(t *testing.T) {
		fx := tearUp(t)

		timeNow := time_utils.TimeNow()

		fx.gameStorage.EXPECT().Insert(fx.ctx, database.Game{
			LeagueID:   1,
			Type:       1,
			Number:     "1",
			PlaceID:    1,
			Date:       timeNow.UTC(),
			Price:      400,
			MaxPlayers: 10,
		}).Return(0, errors.New("some error"))

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID:   int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:       model.GameTypeClassic,
			Number:     "1",
			PlaceID:    1,
			Date:       model.DateTime(timeNow),
			Price:      400,
			MaxPlayers: 10,
		})
		assert.Equal(t, int32(0), got)
		assert.Error(t, err)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		timeNow := time_utils.TimeNow()

		fx.gameStorage.EXPECT().Insert(fx.ctx, database.Game{
			LeagueID:   1,
			Type:       1,
			Number:     "1",
			PlaceID:    1,
			Date:       timeNow.UTC(),
			Price:      400,
			MaxPlayers: 10,
		}).Return(1, nil)

		got, err := fx.facade.AddGame(fx.ctx, model.Game{
			LeagueID:   int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Type:       model.GameTypeClassic,
			Number:     "1",
			PlaceID:    1,
			Date:       model.DateTime(timeNow),
			Price:      400,
			MaxPlayers: 10,
		})
		assert.Equal(t, int32(1), got)
		assert.NoError(t, err)
	})
}
