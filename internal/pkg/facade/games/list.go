package games

import (
	"context"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// ListGames ...
func (f *Facade) ListGames(ctx context.Context) ([]model.Game, error) {
	var modelGames []model.Game
	err := f.db.RunTX(ctx, "ListGames", func(ctx context.Context) error {
		dbGames, err := f.gameStorage.Find(ctx, builder.And(
			builder.IsNull{
				"deleted_at",
			},
		), "date")
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		modelGames = make([]model.Game, 0, len(dbGames))
		for _, dbGame := range dbGames {
			modelGame := convertDBGameToModelGame(dbGame)
			modelGame.HasPassed = !modelGame.IsActive()

			modelGames = append(modelGames, modelGame)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ListGames error: %w", err)
	}

	return modelGames, nil
}
