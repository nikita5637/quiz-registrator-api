package certificatemanager

import (
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

const (
	certificateWonOnGameNotFoundReason   = "WON_ON_GAME_NOT_FOUND"
	certificateSpentOnGameNotFoundReason = "SPENT_ON_GAME_NOT_FOUND"

	minWonOn   = int32(1)
	minSpentOn = int32(1)
)

var (
	certificateNotFoundLexeme = i18n.Lexeme{
		Key:      "certificate_not_found",
		FallBack: "Certificate not found",
	}
	certificateSpentOnGameNotFoundLexeme = i18n.Lexeme{
		Key:      "certificate_spent_on_game_not_found",
		FallBack: "Certificate spent on game not found",
	}
	certificateWonOnGameNotFoundLexeme = i18n.Lexeme{
		Key:      "certificate_won_on_game_not_found",
		FallBack: "Certificate won on game not found",
	}

	errorDetailsByField = map[string]errorDetails{
		"Type": {
			Reason: "INVALID_CERTIFICATE_TYPE",
			Lexeme: i18n.Lexeme{
				Key:      "invalid_certificate_type",
				FallBack: "Invalid certificate type",
			},
		},
		"WonOn": {
			Reason: "INVALID_WON_ON_GAME_ID",
			Lexeme: i18n.Lexeme{
				Key:      "invalid_certificate_won_on_game_id",
				FallBack: "Invalid certificate won on game ID",
			},
		},
		"SpentOn": {
			Reason: "INVALID_SPENT_ON_GAME_ID",
			Lexeme: i18n.Lexeme{
				Key:      "invalid_certificate_spent_on_game_id",
				FallBack: "Invalid certificate spent on game ID",
			},
		},
		"Info": {
			Reason: "INVALID_INFO",
			Lexeme: i18n.Lexeme{
				Key:      "invalid_certificate_info",
				FallBack: "Invalid certificate info",
			},
		},
	}
)

func convertModelCertificateToProtoCertificate(certificate model.Certificate) *certificatemanagerpb.Certificate {
	ret := &certificatemanagerpb.Certificate{
		Id:    certificate.ID,
		Type:  certificatemanagerpb.CertificateType(certificate.Type),
		WonOn: certificate.WonOn,
	}
	if certificate.SpentOn.Valid {
		ret.SpentOn = &wrapperspb.Int32Value{
			Value: certificate.SpentOn.Value,
		}
	}
	if certificate.Info.Valid {
		ret.Info = &wrapperspb.StringValue{
			Value: certificate.Info.Value,
		}
	}

	return ret
}

func validateSpentOn(value interface{}) error {
	v, ok := value.(model.MaybeInt32)
	if !ok {
		return errors.New("must be MaybeInt32")
	}

	return validation.Validate(v.Value, validation.When(v.Valid, validation.Required, validation.Min(minSpentOn)))
}

func validateCertificateInfo(value interface{}) error {
	v, ok := value.(model.MaybeString)
	if !ok {
		return errors.New("must be MaybeString")
	}

	if err := validation.Validate(v.Value, validation.When(v.Valid, validation.Required, validation.Length(1, 256))); err != nil {
		return err
	}

	if valid := json.Valid([]byte(v.Value)); !valid && v.Valid {
		return validation.NewError("validation_invalid_json_value", "must be a valid JSON")
	}

	return nil
}
