package gameresults

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// SearchGameResultByGameID ...
func (f *Facade) SearchGameResultByGameID(ctx context.Context, gameID int32) (model.GameResult, error) {
	var modelGameResult model.GameResult
	err := f.db.RunTX(ctx, "SearchGameResultByGameID", func(ctx context.Context) error {
		var err error
		modelGameResult, err = f.gameResultStorage.GetGameResultByFkGameID(ctx, int(gameID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game result by fk_game_id error: %w", ErrGameResultNotFound)
			}

			return fmt.Errorf("get game result by fk_game_id error: %w", err)
		}

		return nil
	})
	if err != nil {
		return model.GameResult{}, fmt.Errorf("SearchGameResultByGameID error: %w", err)
	}

	return modelGameResult, nil
}
