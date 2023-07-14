package place

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
