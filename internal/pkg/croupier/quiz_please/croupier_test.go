package quiz_please

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := New(Config{
			LotteryLink: "link1",
		})

		assert.Equal(t, &Croupier{
			lotteryLink: "link1",
		}, c)
	})
}
