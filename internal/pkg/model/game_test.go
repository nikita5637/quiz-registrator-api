package model

import (
	"testing"

	"github.com/mono83/maybe"
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
			GameLink:    maybe.Nothing[string](),
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
