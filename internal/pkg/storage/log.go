//go:generate mockery --case underscore --name LogStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// LogStorage ...
type LogStorage interface {
	Create(ctx context.Context, log database.Log) error
}

// NewLogStorage ...
func NewLogStorage(driver string, txManager *tx.Manager) LogStorage {
	switch driver {
	case mysql.DriverName:
		return mysql.NewLogStorageAdapter(txManager)
	}

	return nil
}
