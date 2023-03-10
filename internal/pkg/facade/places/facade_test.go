package places

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ps := storage.NewPlaceStorage(nil)
		cfg := Config{
			PlaceStorage: ps,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			placeStorage: ps,
		}, facade)
	})
}
