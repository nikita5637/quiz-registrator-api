package certificates

import (
	"database/sql"

	"github.com/mono83/maybe"
	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

const (
	// won_on field
	gameIDFK1ConstraintName = "game_id_fk_1"
)

func convertDBCertificateToModelCertificate(certificate database.Certificate) model.Certificate {
	ret := model.Certificate{
		ID:      int32(certificate.ID),
		Type:    model.CertificateType(certificate.Type),
		WonOn:   int32(certificate.WonOn),
		SpentOn: maybe.Nothing[int32](),
		Info:    maybe.Nothing[string](),
	}

	if certificate.SpentOn.Valid {
		ret.SpentOn = maybe.Just(int32(certificate.SpentOn.Int64))
	}

	if certificate.Info.Valid {
		ret.Info = maybe.Just(certificate.Info.String)
	}

	return ret
}

func convertModelCertificateToDBCertificate(certificate model.Certificate) database.Certificate {
	return database.Certificate{
		ID:    int(certificate.ID),
		Type:  certificate.Type.ToSQL(),
		WonOn: int(certificate.WonOn),
		SpentOn: sql.NullInt64{
			Int64: int64(certificate.SpentOn.Value()),
			Valid: certificate.SpentOn.IsPresent(),
		},
		Info: sql.NullString{
			String: certificate.Info.Value(),
			Valid:  certificate.Info.IsPresent(),
		},
	}
}
