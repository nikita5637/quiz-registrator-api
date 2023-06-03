package gameresults

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// ListGameResults ...
func (f *Facade) ListGameResults(ctx context.Context) ([]model.GameResult, error) {
	var modelGameResults []model.GameResult
	err := f.db.RunTX(ctx, "ListGameResults", func(ctx context.Context) error {
		dbGameResults, err := f.gameResultStorage.GetGameResults(ctx)
		if err != nil {
			return fmt.Errorf("get game results error: %w", err)
		}

		modelGameResults = make([]model.GameResult, 0, len(dbGameResults))
		for _, dbGameResult := range dbGameResults {
			modelGameResults = append(modelGameResults, convertDBGameResultToModelGameResult(dbGameResult))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("list game results error: %w", err)
	}

	return modelGameResults, nil
}
