package mysql

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// CertificateStorageAdapter ...
type CertificateStorageAdapter struct {
	certificateStorage *CertificateStorage
}

// NewCertificateStorageAdapter ...
func NewCertificateStorageAdapter(txManager *tx.Manager) *CertificateStorageAdapter {
	return &CertificateStorageAdapter{
		certificateStorage: NewCertificateStorage(txManager),
	}
}

// CreateCertificate ...
func (a *CertificateStorageAdapter) CreateCertificate(ctx context.Context, dbCertificate Certificate) (int32, error) {
	id, err := a.certificateStorage.Insert(ctx, dbCertificate)
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

// DeleteCertificate ...
func (a *CertificateStorageAdapter) DeleteCertificate(ctx context.Context, id int32) error {
	return a.certificateStorage.Delete(ctx, int(id))
}

// GetCertificateByID ...
func (a *CertificateStorageAdapter) GetCertificateByID(ctx context.Context, id int32) (*Certificate, error) {
	return a.certificateStorage.GetCertificateByID(ctx, int(id))
}

// GetCertificates ...
func (a *CertificateStorageAdapter) GetCertificates(ctx context.Context) ([]Certificate, error) {
	return a.certificateStorage.Find(ctx, builder.IsNull{
		"deleted_at",
	}, "id")
}

// PatchCertificate ...
func (a *CertificateStorageAdapter) PatchCertificate(ctx context.Context, dbCertificate Certificate) error {
	return a.certificateStorage.Update(ctx, dbCertificate)
}
