package croupier

import (
	"context"
	"testing"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"

	"github.com/stretchr/testify/assert"
)

func TestCroupier_GetIsLotteryActive(t *testing.T) {
	globalConfig := config.GlobalConfig{}
	globalConfig.LotteryStartsBefore = 3600
	config.UpdateGlobalConfig(globalConfig)

	t.Run("test case 1", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-01-01 15:00")
		}

		ctx := context.Background()
		c := New(Config{})

		game := model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Date:     model.DateTime(timeutils.ConvertTime("2022-01-01 19:00")),
		}

		got := c.GetIsLotteryActive(ctx, game)
		assert.False(t, got)
	})

	t.Run("test case 2", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-01-01 18:00")
		}

		ctx := context.Background()
		c := New(Config{})

		game := model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Date:     model.DateTime(timeutils.ConvertTime("2022-01-01 19:00")),
		}

		got := c.GetIsLotteryActive(ctx, game)
		assert.False(t, got)
	})

	t.Run("test case 3", func(t *testing.T) {
		timeutils.TimeNow = func() time.Time {
			return timeutils.ConvertTime("2022-01-01 18:01")
		}

		ctx := context.Background()
		c := New(Config{})

		game := model.Game{
			LeagueID: int32(leaguepb.LeagueID_QUIZ_PLEASE),
			Date:     model.DateTime(timeutils.ConvertTime("2022-01-01 19:00")),
		}

		got := c.GetIsLotteryActive(ctx, game)
		assert.True(t, got)
	})
}
