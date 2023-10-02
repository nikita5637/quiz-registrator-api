package gameplayers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
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

// GetGamePlayersByFields ...
func (f *Facade) GetGamePlayersByFields(ctx context.Context, gamePlayer model.GamePlayer) ([]model.GamePlayer, error) {
	gamePlayers := make([]model.GamePlayer, 0)
	err := f.db.RunTX(ctx, "GetGamePlayersByFields", func(ctx context.Context) error {
		var dbGamePlayers []database.GamePlayer
		var err error
		builderCond := builder.NewCond().And(
			builder.Eq{
				"fk_game_id":    gamePlayer.GameID,
				"registered_by": gamePlayer.RegisteredBy,
			},
			builder.IsNull{
				"deleted_at",
			},
		)
		if userID, ok := gamePlayer.UserID.Get(); ok {
			dbGamePlayers, err = f.gamePlayerStorage.Find(ctx, builderCond.And(
				builder.Eq{
					"fk_user_id": userID,
				},
			))
		} else {
			dbGamePlayers, err = f.gamePlayerStorage.Find(ctx, builderCond.And(
				builder.IsNull{
					"fk_user_id",
				},
			))
		}
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

// GetUserGameIDs ...
func (f *Facade) GetUserGameIDs(ctx context.Context, userID int32) ([]int32, error) {
	var gameIDs []int32
	err := f.db.RunTX(ctx, "GetUserGameIDs", func(ctx context.Context) error {
		gamePlayers, err := f.gamePlayerStorage.Find(ctx, builder.NewCond().And(
			builder.Eq{
				"fk_user_id": userID,
			},
			builder.IsNull{
				"deleted_at",
			},
		))
		if err != nil {
			return fmt.Errorf("find error: %w", err)
		}

		gameIDs = make([]int32, 0, len(gamePlayers))
		for _, gamePlayer := range gamePlayers {
			gameIDs = append(gameIDs, int32(gamePlayer.FkGameID))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("GetUserGameIDs error: %w", err)
	}

	return gameIDs, nil
}
