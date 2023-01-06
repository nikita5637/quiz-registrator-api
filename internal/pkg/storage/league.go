//go:generate mockery --case underscore --name LeagueStorage --with-expecter

package storage

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// LeagueStorage ...
type LeagueStorage interface {
	GetLeagueByID(ctx context.Context, id int32) (model.League, error)
}

// NewLeagueStorage ...
func NewLeagueStorage(db *sql.DB) LeagueStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewLeagueStorageAdapter(db)
	}

	return nil
}
