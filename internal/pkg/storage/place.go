//go:generate mockery --case underscore --name PlaceStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// PlaceStorage ...
type PlaceStorage interface {
	CreatePlace(ctx context.Context, place database.Place) (int, error)
	GetPlaceByID(ctx context.Context, id int32) (model.Place, error)
}

// NewPlaceStorage ...
func NewPlaceStorage(driver string, txManager *tx.Manager) PlaceStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewPlaceStorageAdapter(txManager)
	}

	return nil
}
