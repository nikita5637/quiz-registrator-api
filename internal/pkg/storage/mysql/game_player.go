package mysql

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GamePlayerStorageAdapter ...
type GamePlayerStorageAdapter struct {
	gamePlayerStorage *GamePlayerStorage
}

// NewGamePlayerStorageAdapter ...
func NewGamePlayerStorageAdapter(db *sql.DB) *GamePlayerStorageAdapter {
	return &GamePlayerStorageAdapter{
		gamePlayerStorage: NewGamePlayerStorage(db),
	}
}

// Delete ...
func (a *GamePlayerStorageAdapter) Delete(ctx context.Context, id int32) error {
	return a.gamePlayerStorage.Delete(ctx, int(id))
}

// Find ...
func (a *GamePlayerStorageAdapter) Find(ctx context.Context, q builder.Cond) ([]model.GamePlayer, error) {
	dbGamePlayers, err := a.gamePlayerStorage.Find(ctx, q, "")
	if err != nil {
		return nil, err
	}

	modelGamePlayers := make([]model.GamePlayer, 0, len(dbGamePlayers))
	for _, dbGamePlayer := range dbGamePlayers {
		modelGamePlayers = append(modelGamePlayers, convertDBGamePlayerToModelGamePlayer(dbGamePlayer))
	}

	return modelGamePlayers, nil
}

// Insert ...
func (a *GamePlayerStorageAdapter) Insert(ctx context.Context, gamePlayer model.GamePlayer) (int32, error) {
	id, err := a.gamePlayerStorage.Insert(ctx, convertModelGamePlayerToDBGamePlayer(gamePlayer))
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

func convertDBGamePlayerToModelGamePlayer(gamePlayer GamePlayer) model.GamePlayer {
	return model.GamePlayer{
		ID:           int32(gamePlayer.ID),
		FkGameID:     int32(gamePlayer.FkGameID),
		FkUserID:     int32(gamePlayer.FkUserID.Int64),
		RegisteredBy: int32(gamePlayer.RegisteredBy),
		Degree:       int32(gamePlayer.Degree),
	}
}

func convertModelGamePlayerToDBGamePlayer(gamePlayer model.GamePlayer) GamePlayer {
	ret := GamePlayer{
		ID:           int(gamePlayer.ID),
		FkGameID:     int(gamePlayer.FkGameID),
		RegisteredBy: int(gamePlayer.RegisteredBy),
		Degree:       uint8(gamePlayer.Degree),
	}

	if gamePlayer.FkUserID > 0 {
		ret.FkUserID = sql.NullInt64{
			Int64: int64(gamePlayer.FkUserID),
			Valid: true,
		}
	}

	return ret
}
