package croupier

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// GetIsLotteryActive ...
func (c *Croupier) GetIsLotteryActive(ctx context.Context, game model.Game) bool {
	lotteryStartsBefore := config.GetValue("LotteryStartsBefore").Uint16()
	if timeutils.TimeNow().UTC().After(game.Date.AsTime().Add(-1 * time.Duration(lotteryStartsBefore) * time.Second)) {
		return true
	}

	return false
}
