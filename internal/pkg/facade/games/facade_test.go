package games

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		gs := storage.NewGameStorage(config.DriverMySQL, nil)
		cfg := Config{
			GameStorage: gs,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			gameStorage: gs,
		}, facade)
	})
}
