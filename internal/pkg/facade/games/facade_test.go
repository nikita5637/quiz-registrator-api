package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := New(Config{})
		assert.NotNil(t, f)
	})
}
