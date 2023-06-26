package certificates

import (
	model "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

const (
	// won_on field
	gameIDFK1ConstraintName = "game_id_fk_1"
)

func convertDBCertificateToModelCertificate(certificate database.Certificate) model.Certificate {
	return model.Certificate{
		ID:      int32(certificate.ID),
		Type:    model.CertificateType(certificate.Type),
		WonOn:   int32(certificate.WonOn),
		SpentOn: model.NewMaybeInt32(int32(certificate.SpentOn.Int64)),
		Info:    model.NewMaybeString(certificate.Info.String),
	}
}

func convertModelCertificateToDBCertificate(certificate model.Certificate) database.Certificate {
	return database.Certificate{
		ID:      int(certificate.ID),
		Type:    certificate.Type.ToSQL(),
		WonOn:   int(certificate.WonOn),
		SpentOn: certificate.SpentOn.ToSQLNullInt64(),
		Info:    certificate.Info.ToSQL(),
	}
}
