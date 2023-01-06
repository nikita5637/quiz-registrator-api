package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// DeleteGame ...
func (f *Facade) DeleteGame(ctx context.Context, gameID int32) error {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("delete game error: %w", model.ErrGameNotFound)
		}

		return fmt.Errorf("delete game error: %w", err)
	}

	if !game.IsActive() {
		return fmt.Errorf("delete game error: %w", model.ErrGameNotFound)
	}

	return f.gameStorage.Delete(ctx, gameID)
}
