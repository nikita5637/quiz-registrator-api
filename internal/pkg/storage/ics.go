//go:generate mockery --case underscore --name ICSFileStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// ICSFileStorage ...
type ICSFileStorage interface {
	CreateICSFile(ctx context.Context, dbICSFile database.IcsFile) (int, error)
	DeleteICSFile(ctx context.Context, id int) error
	GetICSFileByID(ctx context.Context, id int) (*database.IcsFile, error)
	GetICSFileByGameID(ctx context.Context, id int) (*database.IcsFile, error)
	GetICSFiles(ctx context.Context) ([]database.IcsFile, error)
}

// NewICSFileStorage ...
func NewICSFileStorage(txManager *tx.Manager) ICSFileStorage {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewICSFileStorageAdapter(txManager)
	}

	return nil
}
