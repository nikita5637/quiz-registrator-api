package croupier

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// RegisterForLottery ...
func (c *Croupier) RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error) {
	switch game.LeagueID {
	case model.LeagueQuizPlease:
		if !c.GetIsLotteryActive(ctx, game) {
			return 0, fmt.Errorf("quiz please lottery registration error: %w", model.ErrLotteryNotAvailable)
		}

		if user.Email == "" || user.Name == "" || user.Phone == "" {
			return 0, fmt.Errorf("quiz please lottery registration error: %w", model.ErrLotteryPermissionDenied)
		}

		number, err := c.quizPleaseCroupier.RegisterForLottery(ctx, game, user)
		if err != nil {
			return 0, fmt.Errorf("quiz please lottery registration error: %w", err)
		}

		return number, nil
	}

	return 0, fmt.Errorf("register for lottery error: %w", model.ErrLotteryNotImplemented)
}
