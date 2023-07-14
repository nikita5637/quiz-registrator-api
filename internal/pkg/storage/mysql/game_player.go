package mysql

import (
	"context"
	"database/sql"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePlayerStorageAdapter ...
type GamePlayerStorageAdapter struct {
	gamePlayerStorage *GamePlayerStorage
}

// NewGamePlayerStorageAdapter ...
func NewGamePlayerStorageAdapter(txManager *tx.Manager) *GamePlayerStorageAdapter {
	return &GamePlayerStorageAdapter{
		gamePlayerStorage: NewGamePlayerStorage(txManager),
	}
}

// Delete ...
func (a *GamePlayerStorageAdapter) Delete(ctx context.Context, id int32) error {
	return a.gamePlayerStorage.Delete(ctx, int(id))
}

// Find ...
func (a *GamePlayerStorageAdapter) Find(ctx context.Context, q builder.Cond) ([]GamePlayer, error) {
	return a.gamePlayerStorage.Find(ctx, q, "")
}

// Insert ...
func (a *GamePlayerStorageAdapter) Insert(ctx context.Context, gamePlayer model.GamePlayer) (int32, error) {
	id, err := a.gamePlayerStorage.Insert(ctx, convertModelGamePlayerToDBGamePlayer(gamePlayer))
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

func convertModelGamePlayerToDBGamePlayer(gamePlayer model.GamePlayer) GamePlayer {
	ret := GamePlayer{
		ID:       int(gamePlayer.ID),
		FkGameID: int(gamePlayer.FkGameID),
		FkUserID: sql.NullInt64{
			Int64: int64(gamePlayer.FkUserID.Value),
			Valid: gamePlayer.FkUserID.Valid,
		},
		RegisteredBy: int(gamePlayer.RegisteredBy),
		Degree:       uint8(gamePlayer.Degree),
	}

	return ret
}
