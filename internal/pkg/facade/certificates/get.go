package certificates

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GetCertificate ...
func (f *Facade) GetCertificate(ctx context.Context, id int32) (model.Certificate, error) {
	var modelCertificate model.Certificate
	err := f.db.RunTX(ctx, "GetCertificate", func(ctx context.Context) error {
		dbCertificate, err := f.certificateStorage.GetCertificateByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get certificate by ID error: %w", ErrCertificateNotFound)
			}

			return fmt.Errorf("get certificate by ID error: %w", err)
		}

		modelCertificate = convertDBCertificateToModelCertificate(*dbCertificate)
		return nil
	})
	if err != nil {
		return model.Certificate{}, fmt.Errorf("GetCertificate error: %w", err)
	}

	return modelCertificate, nil
}
