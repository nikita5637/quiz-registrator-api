//go:generate mockery --case underscore --name CertificateStorage --with-expecter

package storage

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/config"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// CertificateStorage ...
type CertificateStorage interface {
	CreateCertificate(ctx context.Context, dbCertificate database.Certificate) (int, error)
	DeleteCertificate(ctx context.Context, id int) error
	GetCertificateByID(ctx context.Context, id int) (*database.Certificate, error)
	GetCertificates(ctx context.Context) ([]database.Certificate, error)
	PatchCertificate(ctx context.Context, dbCertificate database.Certificate) error
}

// NewCertificateStorage ...
func NewCertificateStorage(driver string, txManager *tx.Manager) CertificateStorage {
	switch driver {
	case config.DriverMySQL:
		return mysql.NewCertificateStorageAdapter(txManager)
	}

	return nil
}
