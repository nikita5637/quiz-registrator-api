package certificates

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// CreateCertificate ...
func (f *Facade) CreateCertificate(ctx context.Context, certificate model.Certificate) (model.Certificate, error) {
	createdModelCert := model.Certificate{}
	err := f.db.RunTX(ctx, "CreateCertificate", func(ctx context.Context) error {
		newDBCertificate := convertModelCertificateToDBCertificate(certificate)
		id, err := f.certificateStorage.CreateCertificate(ctx, newDBCertificate)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, gameIDFK1ConstraintName); i != -1 {
						return fmt.Errorf("create certificate error: %w", ErrWonOnGameNotFound)
					}

					return fmt.Errorf("create certificate error: %w", ErrSpentOnGameNotFound)
				}
			}

			return fmt.Errorf("create certificate error: %w", err)
		}

		newDBCertificate.ID = id
		createdModelCert = convertDBCertificateToModelCertificate(newDBCertificate)

		return nil
	})
	if err != nil {
		return model.Certificate{}, fmt.Errorf("CreateCertificate error: %w", err)
	}

	return createdModelCert, nil
}
