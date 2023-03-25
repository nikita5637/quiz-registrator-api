package croupier

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
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
	if game.LeagueID == pkgmodel.LeagueSquiz {
		implemented = false
	}

	if !implemented {
		return 0, fmt.Errorf("%w", model.ErrLotteryNotImplemented)
	}

	if !c.GetIsLotteryActive(ctx, game) {
		return 0, fmt.Errorf("%w", model.ErrLotteryNotAvailable)
	}

	if !game.My {
		return 0, fmt.Errorf("%w", model.ErrLotteryPermissionDenied)
	}

	if user.Email == "" || user.Name == "" || user.Phone == "" {
		return 0, fmt.Errorf("%w", model.ErrLotteryPermissionDenied)
	}

	switch game.LeagueID {
	case pkgmodel.LeagueQuizPlease:
		number, err := c.quizPleaseCroupier.RegisterForLottery(ctx, game, user)
		if err != nil {
			return 0, fmt.Errorf("quiz please lottery registration error: %w", err)
		}

		return number, nil
	case pkgmodel.LeagueSquiz:
		_, err := c.squizCroupier.RegisterForLottery(ctx, game, user)
		if err != nil {
			return 0, fmt.Errorf("squiz lottery registration error: %w", err)
		}

		return 0, nil
	}

	return 0, fmt.Errorf("register for lottery error: %w", model.ErrLotteryNotImplemented)
}
