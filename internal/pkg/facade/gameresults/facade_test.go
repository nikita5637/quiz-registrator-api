package gameresults

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		grs := storage.NewGameResultStorage(config.DriverMySQL, nil)
		txManager := tx.NewManager(nil)
		cfg := Config{
			GameResultStorage: grs,
			TxManager:         txManager,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			db:                txManager,
			gameResultStorage: grs,
		}, facade)
	})
}
