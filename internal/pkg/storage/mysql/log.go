package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
)

// LogStorageAdapter ...
type LogStorageAdapter struct {
	logStorage *LogStorage
}

// NewLogStorageAdapter ...
func NewLogStorageAdapter(txManager *tx.Manager) *LogStorageAdapter {
	return &LogStorageAdapter{
		logStorage: NewLogStorage(txManager),
	}
}

// Create ...
func (a *LogStorageAdapter) Create(ctx context.Context, log Log) error {
	log.Timestamp = timeutils.TimeNow().UTC()

	_, err := a.logStorage.Insert(ctx, log)
	return err
}
