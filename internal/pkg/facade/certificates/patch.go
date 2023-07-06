package certificates

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// PatchCertificate ...
func (f *Facade) PatchCertificate(ctx context.Context, certificate model.Certificate) (model.Certificate, error) {
	err := f.db.RunTX(ctx, "PatchCertificate", func(ctx context.Context) error {
		patchedDBCertificate := convertModelCertificateToDBCertificate(certificate)
		if err := f.certificateStorage.PatchCertificate(ctx, patchedDBCertificate); err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, gameIDFK1ConstraintName); i != -1 {
						return fmt.Errorf("patch certificate error: %w", ErrWonOnGameNotFound)
					}

					return fmt.Errorf("patch certificate error: %w", ErrSpentOnGameNotFound)
				}
			}

			return fmt.Errorf("patch certificate error: %w", err)
		}

		return nil

	})
	if err != nil {
		return model.Certificate{}, fmt.Errorf("PatchCertificate error: %w", err)
	}

	return certificate, nil
}
