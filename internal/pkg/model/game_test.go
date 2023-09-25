package model

import (
	"testing"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"

	"github.com/stretchr/testify/assert"
)

func Test_NewGame(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGame()
		assert.Equal(t, Game{
			ExternalID:  maybe.Nothing[int32](),
			Name:        maybe.Nothing[string](),
			PaymentType: maybe.Nothing[string](),
			Payment:     maybe.Nothing[Payment](),
		}, got)
	})
}

func TestGame_DateTime(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var game *Game
		got := game.DateTime()
		assert.Equal(t, DateTime{}, got)
	})

	t.Run("ok", func(t *testing.T) {
		timeNow := timeutils.TimeNow()
		game := &Game{
			Date: DateTime(timeNow),
		}
		got := game.DateTime()
		assert.Equal(t, DateTime(timeNow), got)
	})
}

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

	t.Run("game date + lag less then time_utils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag+1) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game date + lag is equal time_utils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.False(t, got)
	})

	t.Run("game date + lag is greater than time_utils.TimeNow()", func(t *testing.T) {
		gameDate := timeutils.TimeNow()
		timeutils.TimeNow = func() time.Time {
			return gameDate.Add(time.Duration(activeGameLag-1) * time.Second)
		}

		g := &Game{
			Date: DateTime(gameDate),
		}
		got := g.IsActive()
		assert.True(t, got)
	})
}
