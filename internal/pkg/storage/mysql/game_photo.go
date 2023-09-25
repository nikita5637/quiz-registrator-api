package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePhotoStorageAdapter ...
type GamePhotoStorageAdapter struct {
	gamePhotoStorage *GamePhotoStorage
}

// NewGamePhotoStorageAdapter ...
func NewGamePhotoStorageAdapter(txManager *tx.Manager) *GamePhotoStorageAdapter {
	return &GamePhotoStorageAdapter{
		gamePhotoStorage: NewGamePhotoStorage(txManager),
	}
}

// GetGamePhotosByGameID ...
func (a *GamePhotoStorageAdapter) GetGamePhotosByGameID(ctx context.Context, gameID int) ([]*GamePhoto, error) {
	return a.gamePhotoStorage.GetGamePhotoByFkGameID(ctx, gameID)
}

// Insert ...
func (a *GamePhotoStorageAdapter) Insert(ctx context.Context, gamePhoto GamePhoto) (int, error) {
	return a.gamePhotoStorage.Insert(ctx, gamePhoto)
}
