package storage

import (
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// NewDB ...
func NewDB() (*sql.DB, error) {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewDB(config.DriverMySQL, config.GetDatabaseDSN())
	}

	return nil, nil
}
