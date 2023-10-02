package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGamePlayerStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGamePlayerStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}
