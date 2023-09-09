package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGamePhotoStorageAdapter(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		got := NewGamePhotoStorageAdapter(nil)
		assert.NotNil(t, got)
	})
}
