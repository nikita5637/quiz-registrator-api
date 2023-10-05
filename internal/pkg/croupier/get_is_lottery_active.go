package croupier

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
)

// GetIsLotteryActive ...
func (c *Croupier) GetIsLotteryActive(ctx context.Context, game model.Game) bool {
	lotteryStartsBefore := viper.GetDuration("croupier.lottery_starts_before") * time.Second
	if timeutils.TimeNow().UTC().After(game.Date.AsTime().Add(-1 * lotteryStartsBefore)) {
		return true
	}

	return false
}
