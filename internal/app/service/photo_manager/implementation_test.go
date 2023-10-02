//go:generate mockery --case underscore --name GamePhotosFacade --with-expecter

package photomanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		i := New(Config{})
		assert.NotNil(t, i)
	})
}
