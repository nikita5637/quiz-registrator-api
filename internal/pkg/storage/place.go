//go:generate mockery --case underscore --name PlaceStorage --with-expecter

package storage

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// PlaceStorage ...
type PlaceStorage interface {
	GetPlaceByID(ctx context.Context, id int32) (model.Place, error)
}

// NewPlaceStorage ...
func NewPlaceStorage(db *sql.DB) PlaceStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewPlaceStorageAdapter(db)
	}

	return nil
}
