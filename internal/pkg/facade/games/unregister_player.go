package games

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"

	"github.com/go-xorm/builder"
)

// UnregisterPlayer ...
func (f *Facade) UnregisterPlayer(ctx context.Context, gameID, userID, deletedBy int32) (model.UnregisterPlayerStatus, error) {
	game, err := f.gameStorage.GetGameByID(ctx, gameID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UnregisterPlayerStatusInvalid, fmt.Errorf("unregister player error: %w", model.ErrGameNotFound)
		}

		return model.UnregisterPlayerStatusInvalid, fmt.Errorf("unregister player error: %w", err)
	}

	if !game.IsActive() {
		return model.UnregisterPlayerStatusInvalid, fmt.Errorf("unregister player error: %w", model.ErrGameNotFound)
	}

	dbGamePlayers, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().Or(
		builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    gameID,
				"fk_user_id":    userID,
				"registered_by": deletedBy,
			},
			builder.IsNull{
				"deleted_at",
			},
		),
		builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    gameID,
				"registered_by": deletedBy,
			},
			builder.IsNull{
				"deleted_at",
			},
			builder.IsNull{
				"fk_user_id",
			},
		),
	))
	if err != nil {
		return model.UnregisterPlayerStatusInvalid, fmt.Errorf("unregister player error: %w", err)
	}

	if len(dbGamePlayers) == 0 {
		return model.UnregisterPlayerStatusNotRegistered, nil
	}

	for _, dbGamePlayer := range dbGamePlayers {
		if userID == int32(dbGamePlayer.FkUserID.Int64) {
			err = f.gamePlayerStorage.Delete(ctx, int32(dbGamePlayer.ID))
			if err != nil {
				return model.UnregisterPlayerStatusInvalid, fmt.Errorf("unregister player error: %w", err)
			}

			return model.UnregisterPlayerStatusOK, nil
		}
	}

	return model.UnregisterPlayerStatusNotRegistered, nil
}
