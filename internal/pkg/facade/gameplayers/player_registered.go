package gameplayers

import (
	"context"
	"fmt"

	"github.com/go-xorm/builder"
)

// PlayerRegisteredOnGame ...
func (f *Facade) PlayerRegisteredOnGame(ctx context.Context, gameID, userID int32) (bool, error) {
	playerRegistered := false
	err := f.db.RunTX(ctx, "PlayerRegisteredOnGame", func(ctx context.Context) error {
		dbPlayerStorage, err := f.gamePlayerStorage.Find(ctx, builder.And(
			builder.Eq{
				"fk_game_id": gameID,
				"fk_user_id": userID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		if len(dbPlayerStorage) == 1 {
			playerRegistered = true
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return playerRegistered, nil
}
