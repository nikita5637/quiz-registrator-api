package croupier

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetGamesWithActiveLottery ...
func (c *Croupier) GetGamesWithActiveLottery(ctx context.Context) ([]model.Game, error) {
	allGames, err := c.gamesFacade.GetGames(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting games error: %w", err)
	}

	games := make([]model.Game, 0, len(allGames))
	for _, game := range allGames {
		leagueWithLottery := false
		for _, league := range c.leaguesWithLottery {
			if league == game.LeagueID {
				leagueWithLottery = true
				break
			}
		}

		if !leagueWithLottery {
			continue
		}

		if game.IsActive() && c.GetIsLotteryActive(ctx, game) {
			games = append(games, game)
		}
	}

	return games, nil
}
