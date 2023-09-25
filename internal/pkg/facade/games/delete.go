package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteGame ...
func (f *Facade) DeleteGame(ctx context.Context, id int32) error {
	err := f.db.RunTX(ctx, "DeleteGame", func(ctx context.Context) error {
		dbGame, err := f.gameStorage.GetGameByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game by ID error: %w", ErrGameNotFound)
			}

			return fmt.Errorf("get game by ID error: %w", err)
		}

		if dbGame.DeletedAt.Valid {
			return fmt.Errorf("get game by ID error: %w", ErrGameNotFound)
		}

		return f.gameStorage.DeleteGame(ctx, int(id))
	})
	if err != nil {
		return fmt.Errorf("DeleteGame error: %w", err)
	}

	return nil
}
