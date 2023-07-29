package gameplayers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetGamePlayer ...
func (f *Facade) GetGamePlayer(ctx context.Context, id int32) (model.GamePlayer, error) {
	gamePlayer := model.GamePlayer{}
	err := f.db.RunTX(ctx, "GetGamePlayer", func(ctx context.Context) error {
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

		gamePlayer = convertDBGamePlayerToModelGamePlayer(*dbGamePlayer)

		return nil
	})
	if err != nil {
		return model.GamePlayer{}, fmt.Errorf("GetGamePlayer error: %w", err)
	}

	return gamePlayer, nil
}

// GetGamePlayersByGameID ...
func (f *Facade) GetGamePlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error) {
	gamePlayers := make([]model.GamePlayer, 0)
	err := f.db.RunTX(ctx, "GetGamePlayersByGameID", func(ctx context.Context) error {
		dbGamePlayers, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
			builder.Eq{
				"fk_game_id": gameID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		for _, dbGamePlayer := range dbGamePlayers {
			gamePlayers = append(gamePlayers, convertDBGamePlayerToModelGamePlayer(dbGamePlayer))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("GetGamePlayersByGameID error: %w", err)
	}

	return gamePlayers, nil
}
