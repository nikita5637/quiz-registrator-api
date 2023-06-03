package certificates

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

const (
	fieldNameType    = "type"
	fieldNameWonOn   = "won_on"
	fieldNameSpentOn = "spent_on"
	fieldNameInfo    = "info"
)

// PatchCertificate ...
func (f *Facade) PatchCertificate(ctx context.Context, certificate model.Certificate, paths []string) (model.Certificate, error) {
	patchedCert := model.Certificate{}
	err := f.db.RunTX(ctx, "PatchCertificate", func(ctx context.Context) error {
		originalDBCertificate, err := f.certificateStorage.GetCertificateByID(ctx, int(certificate.ID))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("get original certificate error: %w", model.ErrCertificateNotFound)
			}

			return fmt.Errorf("get original certificate error: %w", err)
		}

		patchedDBCertificate := *originalDBCertificate
		for _, path := range paths {
			switch path {
			case fieldNameType:
				patchedDBCertificate.Type = uint8(certificate.Type)
			case fieldNameWonOn:
				patchedDBCertificate.WonOn = int(certificate.WonOn)
			case fieldNameSpentOn:
				patchedDBCertificate.SpentOn = sql.NullInt64{
					Valid: certificate.SpentOn.Valid,
					Int64: int64(certificate.SpentOn.Value),
				}
			case fieldNameInfo:
				patchedDBCertificate.Info = sql.NullString{
					Valid:  certificate.Info.Valid,
					String: certificate.Info.Value,
				}
			}
		}

		err = f.certificateStorage.PatchCertificate(ctx, patchedDBCertificate)
		if err != nil {
			if err, ok := err.(*mysql.MySQLError); ok {
				if err.Number == 1452 {
					if i := strings.Index(err.Message, gameIDFK1ConstraintName); i != -1 {
						return fmt.Errorf("patch certificate error: %w", model.ErrWonOnGameNotFound)
					}

					return fmt.Errorf("patch certificate error: %w", model.ErrSpentOnGameNotFound)
				}
			}

			return fmt.Errorf("patch certificate error: %w", err)
		}

		patchedCert = convertDBCertificateToModelCertificate(patchedDBCertificate)

		return nil

	})
	if err != nil {
		return model.Certificate{}, fmt.Errorf("patch certificate error: %w", err)
	}

	return patchedCert, nil
}
