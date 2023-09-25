//go:generate mockery --case underscore --name GameStorage --with-expecter

package storage

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameStorage ...
type GameStorage interface {
	CreateGame(ctx context.Context, game database.Game) (int, error)
	DeleteGame(ctx context.Context, id int) error
	Find(ctx context.Context, q builder.Cond, sort string) ([]database.Game, error)
	FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]database.Game, error)
	GetGameByID(ctx context.Context, id int) (*database.Game, error)
	PatchGame(ctx context.Context, game database.Game) error
	Total(ctx context.Context, q builder.Cond) (uint64, error)
}

// NewGameStorage ...
func NewGameStorage(driver string, txManager *tx.Manager) GameStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewGameStorageAdapter(txManager)
	}

	return nil
}
