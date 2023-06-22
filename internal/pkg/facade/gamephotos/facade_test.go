package gamephotos

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		gps := storage.NewGamePhotoStorage("driver", nil)
		txManager := tx.NewManager(nil)
		cfg := Config{
			GamePhotoStorage: gps,
			TxManager:        txManager,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			db:               txManager,
			gamePhotoStorage: gps,
		}, facade)
	})
}
