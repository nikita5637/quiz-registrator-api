package certificates

import (
	"context"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// ListCertificates ...
func (f *Facade) ListCertificates(ctx context.Context) ([]model.Certificate, error) {
	var modelCertificates []model.Certificate
	err := f.db.RunTX(ctx, "ListCertificates", func(ctx context.Context) error {
		dbCertificates, err := f.certificateStorage.GetCertificates(ctx)
		if err != nil {
			return fmt.Errorf("get certificates error: %w", err)
		}

		modelCertificates = make([]model.Certificate, 0, len(dbCertificates))
		for _, dbCertificate := range dbCertificates {
			modelCertificates = append(modelCertificates, convertDBCertificateToModelCertificate(dbCertificate))
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ListCertificates error: %w", err)
	}

	return modelCertificates, nil
}
