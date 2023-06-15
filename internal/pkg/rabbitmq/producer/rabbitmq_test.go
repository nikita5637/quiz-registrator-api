//go:generate mockery --case underscore --name Channel --with-expecter

package rabbitmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		p := New(Config{})
		assert.NotNil(t, p)
	})
}
