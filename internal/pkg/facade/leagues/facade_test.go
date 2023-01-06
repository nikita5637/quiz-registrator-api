package leagues

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ls := storage.NewLeagueStorage(nil)
		cfg := Config{
			LeagueStorage: ls,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			leagueStorage: ls,
		}, facade)
	})
}
