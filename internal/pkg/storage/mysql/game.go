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

// Delete ...
func (a *GameStorageAdapter) Delete(ctx context.Context, id int) error {
	return a.gameStorage.Delete(ctx, id)
}

// Find ...
func (a *GameStorageAdapter) Find(ctx context.Context, q builder.Cond, sort string) ([]Game, error) {
	return a.gameStorage.Find(ctx, q, sort)
}

// GetGameByID ...
func (a *GameStorageAdapter) GetGameByID(ctx context.Context, id int) (*Game, error) {
	return a.gameStorage.GetGameByID(ctx, id)
}

// Insert ...
func (a *GameStorageAdapter) Insert(ctx context.Context, game Game) (int, error) {
	return a.gameStorage.Insert(ctx, game)
}

// Update ....
func (a *GameStorageAdapter) Update(ctx context.Context, game Game) error {
	return a.gameStorage.Update(ctx, game)
}
