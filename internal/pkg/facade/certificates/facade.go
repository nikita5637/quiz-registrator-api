package certificates

import (
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// Facade ...
type Facade struct {
	db                 *tx.Manager
	certificateStorage storage.CertificateStorage
}

// Config ...
type Config struct {
	CertificateStorage storage.CertificateStorage
	TxManager          *tx.Manager
}

// NewFacade ...
func NewFacade(cfg Config) *Facade {
	return &Facade{
		db:                 cfg.TxManager,
		certificateStorage: cfg.CertificateStorage,
	}
}
