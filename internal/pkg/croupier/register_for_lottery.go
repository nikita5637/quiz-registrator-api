package croupier

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

// RegisterForLottery ...
func (c *Croupier) RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error) {
	implemented := false
	for _, leagueWithLottery := range c.leaguesWithLottery {
		if game.LeagueID == leagueWithLottery {
			implemented = true
			break
		}
	}

	// TODO fix squiz lottery registration
	if game.LeagueID == int32(leaguepb.LeagueID_SQUIZ) {
		implemented = false
	}

	if !implemented {
		return 0, fmt.Errorf("%w", ErrLotteryNotImplemented)
	}

	if !c.GetIsLotteryActive(ctx, game) {
		return 0, fmt.Errorf("%w", ErrLotteryNotAvailable)
	}

	switch game.LeagueID {
	case int32(leaguepb.LeagueID_QUIZ_PLEASE):
		number, err := c.quizPleaseCroupier.RegisterForLottery(ctx, game, user)
		if err != nil {
			return 0, fmt.Errorf("quiz please lottery registration error: %w", err)
		}

		return number, nil
	case int32(leaguepb.LeagueID_SQUIZ):
		_, err := c.squizCroupier.RegisterForLottery(ctx, game, user)
		if err != nil {
			return 0, fmt.Errorf("squiz lottery registration error: %w", err)
		}

		return 0, nil
	}

	return 0, fmt.Errorf("register for lottery error: %w", ErrLotteryNotImplemented)
}
