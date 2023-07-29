package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteGame ...
func (f *Facade) DeleteGame(ctx context.Context, gameID int32) error {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("delete game error: %w", ErrGameNotFound)
		}

		return fmt.Errorf("delete game error: %w", err)
	}

	if !game.IsActive() {
		return fmt.Errorf("delete game error: %w", ErrGameNotFound)
	}

	return f.gameStorage.Delete(ctx, gameID)
}
