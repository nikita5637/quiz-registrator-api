//go:generate mockery --case underscore --name GameResultStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameResultStorage ...
type GameResultStorage interface {
	CreateGameResult(ctx context.Context, dbGameResult database.GameResult) (int, error)
	GetGameResultByID(ctx context.Context, id int) (*database.GameResult, error)
	GetGameResults(ctx context.Context) ([]database.GameResult, error)
	GetGameResultByFkGameID(ctx context.Context, fkGameID int) (model.GameResult, error)
	PatchGameResult(ctx context.Context, dbGameResult database.GameResult) error
}

// NewGameResultStorage ...
func NewGameResultStorage(driver string, txManager *tx.Manager) GameResultStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewGameResultStorageAdapter(txManager)
	}

	return nil
}
