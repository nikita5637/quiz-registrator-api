package quizlogger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		l := New(Config{})
		assert.NotNil(t, l)
	})
}
