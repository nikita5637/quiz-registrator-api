package certificatemanager

import (
	"encoding/json"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
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
	reasonCertificateWonOnGameNotFound   = "WON_ON_GAME_NOT_FOUND"
	reasonCertificateSpentOnGameNotFound = "SPENT_ON_GAME_NOT_FOUND"
	reasonInvalidCertificateType         = "INVALID_CERTIFICATE_TYPE"
	reasonInvalidSpentOnGameID           = "INVALID_SPENT_ON_GAME_ID"
	reasonInvalidWonOnGameID             = "INVALID_WON_ON_GAME_ID"
	reasonInvalidInfo                    = "INVALID_INFO"

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
			Reason: reasonInvalidCertificateType,
			Lexeme: invalidCertificateTypeLexeme,
		},
		"WonOn": {
			Reason: reasonInvalidWonOnGameID,
			Lexeme: invalidWonOnGameIDLexeme,
		},
		"SpentOn": {
			Reason: reasonInvalidSpentOnGameID,
			Lexeme: invalidSpentOnGameIDLexeme,
		},
		"Info": {
			Reason: reasonInvalidInfo,
			Lexeme: invalidInfoLexeme,
		},
	}

	invalidCertificateTypeLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_type",
		FallBack: "Invalid certificate type",
	}
	invalidInfoLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_info",
		FallBack: "Invalid certificate info",
	}
	invalidSpentOnGameIDLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_spent_on_game_id",
		FallBack: "Invalid certificate spent on game ID",
	}
	invalidWonOnGameIDLexeme = i18n.Lexeme{
		Key:      "invalid_certificate_won_on_game_id",
		FallBack: "Invalid certificate won on game ID",
	}
)

func convertModelCertificateToProtoCertificate(certificate model.Certificate) *certificatemanagerpb.Certificate {
	ret := &certificatemanagerpb.Certificate{
		Id:    certificate.ID,
		Type:  certificatemanagerpb.CertificateType(certificate.Type),
		WonOn: certificate.WonOn,
	}
	if v, ok := certificate.SpentOn.Get(); ok {
		ret.SpentOn = &wrapperspb.Int32Value{
			Value: v,
		}
	}
	if v, ok := certificate.Info.Get(); ok {
		ret.Info = &wrapperspb.StringValue{
			Value: v,
		}
	}

	return ret
}

func convertProtoCertificateToModelCertificate(certificate *certificatemanagerpb.Certificate) model.Certificate {
	ret := model.Certificate{
		ID:      certificate.GetId(),
		Type:    model.CertificateType(certificate.GetType()),
		WonOn:   certificate.GetWonOn(),
		SpentOn: maybe.Nothing[int32](),
		Info:    maybe.Nothing[string](),
	}

	if certificate.GetSpentOn() != nil {
		ret.SpentOn = maybe.Just(certificate.GetSpentOn().GetValue())
	}

	if certificate.GetInfo() != nil {
		ret.Info = maybe.Just(certificate.GetInfo().GetValue())
	}

	return ret
}

func getErrorDetails(keys []string) *errorDetails {
	if len(keys) == 0 {
		return nil
	}

	if v, ok := errorDetailsByField[keys[0]]; ok {
		return &v
	}

	return nil
}

func validateSpentOn(value interface{}) error {
	v, ok := value.(maybe.Maybe[int32])
	if !ok {
		return errors.New("must be Maybe[int32]")
	}

	return validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Min(minSpentOn)))
}

func validateCertificateInfo(value interface{}) error {
	v, ok := value.(maybe.Maybe[string])
	if !ok {
		return errors.New("must be Maybe[string]")
	}

	if err := validation.Validate(v.Value(), validation.When(v.IsPresent(), validation.Required, validation.Length(1, 256))); err != nil {
		return err
	}

	if valid := json.Valid([]byte(v.Value())); !valid && v.IsPresent() {
		return validation.NewError("validation_invalid_json_value", "must be a valid JSON")
	}

	return nil
}
