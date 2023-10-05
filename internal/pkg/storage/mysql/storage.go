package mysql

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

const (
	// DriverName ...
	DriverName = "mysql"
)

// NewDB ...
func NewDB(ctx context.Context, dataSourceName string) (*sql.DB, error) {
	logger.DebugKV(ctx, "initialize database connection started", "driverName", DriverName, "DSN", dataSourceName)
	db, err := sql.Open(DriverName, dataSourceName)
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
