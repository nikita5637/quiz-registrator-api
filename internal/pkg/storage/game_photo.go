//go:generate mockery --case underscore --name GamePhotoStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePhotoStorage ...
type GamePhotoStorage interface {
	GetGameIDsWithPhotos(ctx context.Context, limit uint32) ([]int32, error)
	GetGamePhotosByGameID(ctx context.Context, gameID int32) ([]model.GamePhoto, error)
	Insert(ctx context.Context, gamePhoto model.GamePhoto) (int32, error)
}

// NewGamePhotoStorage ...
func NewGamePhotoStorage(txManager *tx.Manager) GamePhotoStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGamePhotoStorageAdapter(txManager)
	}

	return nil
}
