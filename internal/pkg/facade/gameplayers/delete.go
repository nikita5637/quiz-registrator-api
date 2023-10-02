package gameplayers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteGamePlayer ...
func (f *Facade) DeleteGamePlayer(ctx context.Context, id int32) error {
	err := f.db.RunTX(ctx, "DeleteGamePlayer", func(ctx context.Context) error {
		dbGamePlayer, err := f.gamePlayerStorage.GetGamePlayer(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get game player error: %w", ErrGamePlayerNotFound)
			}

			return fmt.Errorf("get game player error: %w", err)
		}

		if dbGamePlayer.DeletedAt.Valid {
			return fmt.Errorf("get game player error: %w", ErrGamePlayerNotFound)
		}

		return f.gamePlayerStorage.DeleteGamePlayer(ctx, int(id))
	})
	if err != nil {
		return fmt.Errorf("DeleteGamePlayer error: %w", err)
	}

	return nil

}
