//go:generate mockery --case underscore --name CertificatesFacade --with-expecter

package certificatemanager

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
)

// CertificatesFacade ...
type CertificatesFacade interface {
	CreateCertificate(ctx context.Context, certificate model.Certificate) (model.Certificate, error)
	DeleteCertificate(ctx context.Context, id int32) error
	GetCertificate(ctx context.Context, id int32) (model.Certificate, error)
	ListCertificates(ctx context.Context) ([]model.Certificate, error)
	PatchCertificate(ctx context.Context, certificate model.Certificate) (model.Certificate, error)
}

// CertificateManager ...
type CertificateManager struct {
	certificatesFacade CertificatesFacade

	certificatemanagerpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	CertificatesFacade CertificatesFacade
}

// New ...
func New(cfg Config) *CertificateManager {
	return &CertificateManager{
		certificatesFacade: cfg.CertificatesFacade,
	}
}
