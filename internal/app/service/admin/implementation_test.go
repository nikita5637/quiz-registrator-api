package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		a := New(Config{})
		assert.NotNil(t, a)
	})
}
