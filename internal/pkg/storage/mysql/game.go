package mysql

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameStorageAdapter ...
type GameStorageAdapter struct {
	gameStorage *GameStorage
}

// NewGameStorageAdapter ...
func NewGameStorageAdapter(txManager *tx.Manager) *GameStorageAdapter {
	return &GameStorageAdapter{
		gameStorage: NewGameStorage(txManager),
	}
}

// CreateGame ...
func (a *GameStorageAdapter) CreateGame(ctx context.Context, game Game) (int, error) {
	return a.gameStorage.Insert(ctx, game)
}

// DeleteGame ...
func (a *GameStorageAdapter) DeleteGame(ctx context.Context, id int) error {
	return a.gameStorage.Delete(ctx, id)
}

// Find ...
func (a *GameStorageAdapter) Find(ctx context.Context, q builder.Cond, sort string) ([]Game, error) {
	return a.gameStorage.Find(ctx, q, sort)
}

// FindWithLimit ...
func (a *GameStorageAdapter) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]Game, error) {
	return a.gameStorage.FindWithLimit(ctx, q, sort, offset, limit)
}

// GetGameByID ...
func (a *GameStorageAdapter) GetGameByID(ctx context.Context, id int) (*Game, error) {
	return a.gameStorage.GetGameByID(ctx, id)
}

// PatchGame ....
func (a *GameStorageAdapter) PatchGame(ctx context.Context, game Game) error {
	return a.gameStorage.Update(ctx, game)
}

// Total ...
func (a *GameStorageAdapter) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	return a.gameStorage.Total(ctx, q)
}
