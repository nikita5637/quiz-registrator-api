//go:generate mockery --case underscore --name UsersFacade --with-expecter

package authentication

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := New(Config{})
		assert.NotNil(t, m)
	})
}
