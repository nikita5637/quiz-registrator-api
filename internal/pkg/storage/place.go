//go:generate mockery --case underscore --name PlaceStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// PlaceStorage ...
type PlaceStorage interface {
	GetPlaceByID(ctx context.Context, id int32) (model.Place, error)
}

// NewPlaceStorage ...
func NewPlaceStorage(txManager *tx.Manager) PlaceStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewPlaceStorageAdapter(txManager)
	}

	return nil
}
