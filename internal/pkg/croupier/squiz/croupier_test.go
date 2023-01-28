package squiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := New(Config{
			LotteryInfoPageLink:     "link1",
			LotteryRegistrationLink: "link2",
		})

		assert.Equal(t, &Croupier{
			lotteryInfoPageLink:     "link1",
			lotteryRegistrationLink: "link2",
		}, c)
	})
}
