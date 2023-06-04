package model

import (
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"

	"github.com/stretchr/testify/assert"
)

func TestGame_IsActive(t *testing.T) {
	activeGameLag := uint16(3600)
	assert.Greater(t, activeGameLag, uint16(1))

	cfg := config.GlobalConfig{}
	cfg.ActiveGameLag = activeGameLag

	config.UpdateGlobalConfig(cfg)

	t.Run("game is nil", func(t *testing.T) {
		var g *Game
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game deleted at is not zero", func(t *testing.T) {
		g := &Game{}
		g.DeletedAt = DateTime(time_utils.TimeNow())
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game date + lag less then time_utils.TimeNow()", func(t *testing.T) {
		gameDate := time_utils.TimeNow()
		time_utils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag+1) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game date + lag is equal time_utils.TimeNow()", func(t *testing.T) {
		gameDate := time_utils.TimeNow()
		time_utils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game date + lag is greater than time_utils.TimeNow()", func(t *testing.T) {
		gameDate := time_utils.TimeNow()
		time_utils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag-1) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.True(t, got)
	})
}

func TestFacade_ValidateGame(t *testing.T) {
	t.Run("league id validation error", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: -1,
			Number:   "1",
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidLeagueID)
	})

	t.Run("game type validation error case 1", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     0,
			Number:   "1",
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidGameType)
	})

	t.Run("game type validation error case 2", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     3,
			Number:   "1",
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidGameType)
	})

	t.Run("game number validation error. game type is classic", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     pkgmodel.GameTypeClassic,
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidGameNumber)
	})

	t.Run("place id validation error", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     pkgmodel.GameTypeClassic,
			Number:   "1",
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidPlaceID)
	})

	t.Run("date validation error", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     pkgmodel.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidDate)
	})

	t.Run("price validation error", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     pkgmodel.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
			Date:     DateTime(time_utils.TimeNow()),
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidPrice)
	})

	t.Run("max players validation error", func(t *testing.T) {
		err := ValidateGame(Game{
			LeagueID: pkgmodel.LeagueQuizPlease,
			Type:     pkgmodel.GameTypeClassic,
			Number:   "1",
			PlaceID:  1,
			Date:     DateTime(time_utils.TimeNow()),
			Price:    400,
		})
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrInvalidMaxPlayers)
	})

	t.Run("ok. game type is classic", func(t *testing.T) {
		timeNow := time_utils.TimeNow()

		err := ValidateGame(Game{
			LeagueID:   pkgmodel.LeagueQuizPlease,
			Type:       pkgmodel.GameTypeClassic,
			Number:     "1",
			PlaceID:    1,
			Date:       DateTime(timeNow),
			Price:      400,
			MaxPlayers: 10,
		})
		assert.NoError(t, err)
	})

	t.Run("ok. game type is closed", func(t *testing.T) {
		timeNow := time_utils.TimeNow()

		err := ValidateGame(Game{
			LeagueID:   pkgmodel.LeagueQuizPlease,
			Type:       pkgmodel.GameTypeClosed,
			Number:     "",
			PlaceID:    1,
			Date:       DateTime(timeNow),
			Price:      400,
			MaxPlayers: 10,
		})
		assert.NoError(t, err)
	})
}
