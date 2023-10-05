//go:generate mockery --case underscore --name LeagueStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// LeagueStorage ...
type LeagueStorage interface {
	CreateLeague(ctx context.Context, league database.League) (int, error)
	GetLeagueByID(ctx context.Context, id int32) (model.League, error)
}

// NewLeagueStorage ...
func NewLeagueStorage(driver string, txManager *tx.Manager) LeagueStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewLeagueStorageAdapter(txManager)
	}

	return nil
}
