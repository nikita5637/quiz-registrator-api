package users

import (
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/stretchr/testify/assert"
)

func TestNewFacade(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		us := storage.NewUserStorage(mysql.DriverName, nil)
		cfg := Config{
			UserStorage: us,
		}
		facade := NewFacade(cfg)

		assert.NotNil(t, facade)
		assert.Equal(t, &Facade{
			userStorage: us,
		}, facade)
	})
}
