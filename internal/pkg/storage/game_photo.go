//go:generate mockery --case underscore --name GamePhotoStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePhotoStorage ...
type GamePhotoStorage interface {
	GetGamePhotosByGameID(ctx context.Context, gameID int) ([]*database.GamePhoto, error)
	Insert(ctx context.Context, gamePhoto database.GamePhoto) (int, error)
}

// NewGamePhotoStorage ...
func NewGamePhotoStorage(driver string, txManager *tx.Manager) GamePhotoStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewGamePhotoStorageAdapter(txManager)
	}

	return nil
}
