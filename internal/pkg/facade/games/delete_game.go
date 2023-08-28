package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteGame ...
func (f *Facade) DeleteGame(ctx context.Context, id int32) error {
	dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("delete game error: %w", ErrGameNotFound)
		}

		return fmt.Errorf("delete game error: %w", err)
	}

	modelGame := convertDBGameToModelGame(*dbGame)

	if !modelGame.IsActive() {
		return fmt.Errorf("delete game error: %w", ErrGameNotFound)
	}

	return f.gameStorage.Delete(ctx, int(id))
}
