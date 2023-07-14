package certificates

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// DeleteCertificate ...
func (f *Facade) DeleteCertificate(ctx context.Context, id int32) error {
	err := f.db.RunTX(ctx, "DeleteCertificate", func(ctx context.Context) error {
		dbCertificate, err := f.certificateStorage.GetCertificateByID(ctx, int(id))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get certificate by ID error: %w", ErrCertificateNotFound)
			}

			return fmt.Errorf("get certificate by ID error: %w", err)
		}

		if dbCertificate.DeletedAt.Valid {
			return fmt.Errorf("get certificate by ID error: %w", ErrCertificateNotFound)
		}

		return f.certificateStorage.DeleteCertificate(ctx, int(id))
	})
	if err != nil {
		return fmt.Errorf("DeleteCertificate error: %w", err)
	}

	return nil
}
