package games

import (
	"context"
	"fmt"
	"time"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/spf13/viper"
)

// SearchGamesByLeagueID ...
func (f *Facade) SearchGamesByLeagueID(ctx context.Context, leagueID int32, offset, limit uint64) ([]model.Game, uint64, error) {
	var modelGames []model.Game
	var total uint64
	err := f.db.RunTX(ctx, "SearchGamesByLeagueID", func(ctx context.Context) error {
		var err error
		total, err = f.gameStorage.Total(ctx, builder.And(
			builder.Eq{
				"league_id": leagueID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("get total error: %w", err)
		}

		dbGames, err := f.gameStorage.FindWithLimit(ctx, builder.And(
			builder.Eq{
				"league_id": leagueID,
			},
			builder.IsNull{
				"deleted_at",
			},
		), "date", offset, limit)
		if err != nil {
			return fmt.Errorf("find with limit error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = gameHasPassed(modelGame)

			modelGames = append(modelGames, modelGame)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("SearchGamesByLeagueID error: %w", err)
	}

	return modelGames, total, nil
}

// SearchPassedAndRegisteredGames ...
func (f *Facade) SearchPassedAndRegisteredGames(ctx context.Context, offset, limit uint64) ([]model.Game, uint64, error) {
	var modelGames []model.Game
	var total uint64
	err := f.db.RunTX(ctx, "SearchPassedAndRegisteredGames", func(ctx context.Context) error {
		hasPassedGameLag := viper.GetDuration("service.game.has_passed_game_lag") * time.Second
		var err error
		total, err = f.gameStorage.Total(ctx, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().UTC().Add(-1 * hasPassedGameLag),
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("total error: %w", err)
		}

		dbGames, err := f.gameStorage.FindWithLimit(ctx, builder.And(
			builder.Eq{
				"registered": true,
			},
			builder.Lt{
				"date": timeutils.TimeNow().UTC().Add(-1 * hasPassedGameLag),
			},
			builder.IsNull{
				"deleted_at",
			},
		), "-date", offset, limit)
		if err != nil {
			return fmt.Errorf("find with limit error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = true

			modelGames = append(modelGames, modelGame)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("SearchPassedAndRegisteredGames error: %w", err)
	}

	return modelGames, total, nil
}