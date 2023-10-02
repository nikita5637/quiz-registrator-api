package mysql

import (
	"context"

	"github.com/go-xorm/builder"
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

// CreateGamePlayer ...
func (a *GamePlayerStorageAdapter) CreateGamePlayer(ctx context.Context, gamePlayer GamePlayer) (int, error) {
	return a.gamePlayerStorage.Insert(ctx, gamePlayer)
}

// DeleteGamePlayer ...
func (a *GamePlayerStorageAdapter) DeleteGamePlayer(ctx context.Context, id int) error {
	return a.gamePlayerStorage.Delete(ctx, id)
}

// Find ...
func (a *GamePlayerStorageAdapter) Find(ctx context.Context, q builder.Cond) ([]GamePlayer, error) {
	return a.gamePlayerStorage.Find(ctx, q, "")
}

// GetGamePlayer ...
func (a *GamePlayerStorageAdapter) GetGamePlayer(ctx context.Context, id int) (*GamePlayer, error) {
	return a.gamePlayerStorage.GetGamePlayerByID(ctx, id)
}

// PatchGamePlayer ...
func (a *GamePlayerStorageAdapter) PatchGamePlayer(ctx context.Context, gamePlayer GamePlayer) error {
	return a.gamePlayerStorage.Update(ctx, gamePlayer)
}
