//go:generate mockery --case underscore --name GameResultStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameResultStorage ...
type GameResultStorage interface {
	GetGameResultByFkGameID(ctx context.Context, fkGameID int32) (model.GameResult, error)
}

// NewGameResultStorage ...
func NewGameResultStorage(txManager *tx.Manager) GameResultStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewGameResultStorageAdapter(txManager)
	}

	return nil
}
