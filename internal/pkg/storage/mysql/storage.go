package mysql

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

// NewDB ...
func NewDB(driverName, dataSourceName string) (*sql.DB, error) {
	ctx := context.Background()
	logger.DebugKV(ctx, "initialize database connection started", "driverName", driverName, "DSN", dataSourceName)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	logger.Debug(ctx, "initialize database connection done")
	return db, nil
}
