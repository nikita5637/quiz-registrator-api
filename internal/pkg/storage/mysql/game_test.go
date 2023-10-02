package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGameStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}
