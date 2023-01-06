package model

import (
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
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
