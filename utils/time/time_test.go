package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConvertTime(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		got := ConvertTime("invalid time")
		assert.Equal(t, time.Time{}, got)
	})

	t.Run("test case 2", func(t *testing.T) {
		got := ConvertTime("2022-03-21 19:30")
		expect, err := time.Parse(time.RFC3339, "2022-03-21T19:30:00Z")
		assert.NoError(t, err)
		assert.Equal(t, expect, got)
	})
}
