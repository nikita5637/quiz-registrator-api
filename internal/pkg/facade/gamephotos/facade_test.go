package gamephotos

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		gps := storage.NewGamePhotoStorage(nil)
		cfg := Config{
			GamePhotoStorage: gps,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			gamePhotoStorage: gps,
		}, facade)
	})
}
