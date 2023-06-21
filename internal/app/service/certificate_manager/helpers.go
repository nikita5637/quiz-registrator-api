package certificatemanager

import (
	"context"
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
)

var (
	errInvalidJSONInfoValue   = errors.New("invalid JSON info value")
	errInvalidCertificateType = errors.New("invalid certificate type")

	certificateNotFoundLexeme = i18n.Lexeme{
		Key:      "certificate_not_found",
		FallBack: "Certificate not found",
	}
	invalidCertificateInfoJSONValueLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_info_json_value",
		FallBack: "Invalid certificate info JSON value",
	}
	invalidCertificateTypeLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_type",
		FallBack: "Invalid certificate type",
	}
	spentOnGameNotFoundLexeme = i18n.Lexeme{
		Key:      "spent_on_game_not_found",
		FallBack: "Spent on game not found",
	}
	wonOnGameNotFoundLexeme = i18n.Lexeme{
		Key:      "won_on_game_not_found",
		FallBack: "Won on game not found",
	}
)

func convertModelCertificateToProtoCertificate(certificate model.Certificate) *certificatemanagerpb.Certificate {
	ret := &certificatemanagerpb.Certificate{
		Id:      certificate.ID,
		Type:    certificatemanagerpb.CertificateType(certificate.Type),
		WonOn:   certificate.WonOn,
		SpentOn: certificate.SpentOn.Value,
		Info:    certificate.Info.Value,
	}

	return ret
}

func validateCertificate(ctx context.Context, certificate *certificatemanagerpb.Certificate) error {
	if valid := json.Valid([]byte(certificate.GetInfo())); !valid {
		return errInvalidJSONInfoValue
	}

	err := validation.Validate(certificate.GetType(), validation.Required, validation.Min(int32(1)), validation.Max(int32(pkgmodel.NumberOfCertificateTypes-1)))
	if err != nil {
		return errInvalidCertificateType
	}

	if len(certificate.GetInfo()) > 256 {
		return errInvalidJSONInfoValue
	}

	return nil
}
