//go:generate mockery --case underscore --name GamePlayerStorage --with-expecter

package storage

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GamePlayerStorage ...
type GamePlayerStorage interface {
	CreateGamePlayer(ctx context.Context, gamePlayer database.GamePlayer) (int, error)
	DeleteGamePlayer(ctx context.Context, id int) error
	Find(ctx context.Context, q builder.Cond) ([]database.GamePlayer, error)
	GetGamePlayer(ctx context.Context, id int) (*database.GamePlayer, error)
	PatchGamePlayer(ctx context.Context, gamePlayer database.GamePlayer) error
}

// NewGamePlayerStorage ...
func NewGamePlayerStorage(driver string, txManager *tx.Manager) GamePlayerStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewGamePlayerStorageAdapter(txManager)
	}

	return nil
}
