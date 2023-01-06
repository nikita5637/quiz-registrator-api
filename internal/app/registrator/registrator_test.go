package registrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegistrator(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		r := New(Config{})
		assert.NotNil(t, r)
	})
}
