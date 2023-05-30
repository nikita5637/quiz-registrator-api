package certificates

import (
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
)

const (
	// won_on field
	gameIDFK1ConstraintName = "game_id_fk_1"
)

func convertDBCertificateToModelCertificate(dbCertificate database.Certificate) model.Certificate {
	return model.Certificate{
		ID:    int32(dbCertificate.ID),
		Type:  pkgmodel.CertificateType(dbCertificate.Type),
		WonOn: int32(dbCertificate.WonOn),
		SpentOn: model.MaybeInt32{
			Valid: dbCertificate.SpentOn.Valid,
			Value: int32(dbCertificate.SpentOn.Int64),
		},
		Info: model.MaybeString{
			Valid: dbCertificate.Info.Valid,
			Value: dbCertificate.Info.String,
		},
	}
}

func convertModelCertificateToDBCertificate(modelCertificate model.Certificate) database.Certificate {
	return database.Certificate{
		ID:    int(modelCertificate.ID),
		Type:  uint8(modelCertificate.Type),
		WonOn: int(modelCertificate.WonOn),
		SpentOn: sql.NullInt64{
			Valid: modelCertificate.SpentOn.Valid,
			Int64: int64(modelCertificate.SpentOn.Value),
		},
		Info: sql.NullString{
			Valid:  modelCertificate.Info.Valid,
			String: modelCertificate.Info.Value,
		},
	}
}
