package certificatemanager

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorDetails struct {
	Reason string
	Lexeme i18n.Lexeme
}

// CreateCertificate ...
func (m *CertificateManager) CreateCertificate(ctx context.Context, req *certificatemanagerpb.CreateCertificateRequest) (*certificatemanagerpb.Certificate, error) {
	createdCertificate := model.Certificate{
		Type:  model.CertificateType(req.GetCertificate().GetType()),
		WonOn: req.GetCertificate().GetWonOn(),
		SpentOn: model.MaybeInt32{
			Valid: req.GetCertificate().GetSpentOn() != nil,
			Value: req.GetCertificate().GetSpentOn().GetValue(),
		},
		Info: model.MaybeString{
			Valid: req.GetCertificate().GetInfo() != nil,
			Value: req.GetCertificate().GetInfo().GetValue(),
		},
	}
	if err := validateCreatedCertificate(createdCertificate); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if validationErrors, ok := err.(validation.Errors); ok && len(validationErrors) > 0 {
			keys := make([]string, 0, len(validationErrors))
			for k := range validationErrors {
				keys = append(keys, k)
			}

			if ed, ok := errorDetailsByField[keys[0]]; ok {
				st = status.New(codes.InvalidArgument, fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()))
				errorInfo := &errdetails.ErrorInfo{
					Reason: ed.Reason,
					Metadata: map[string]string{
						"error": err.Error(),
					},
				}
				localizedMessage := &errdetails.LocalizedMessage{
					Locale:  i18n.GetLangFromContext(ctx),
					Message: i18n.GetTranslator(ed.Lexeme)(ctx),
				}
				st, _ = st.WithDetails(errorInfo, localizedMessage)
			}
		}

		return nil, st.Err()
	}

	certificate, err := m.certificatesFacade.CreateCertificate(ctx, createdCertificate)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, certificates.ErrWonOnGameNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err, certificateWonOnGameNotFoundReason, certificateWonOnGameNotFoundLexeme)
		} else if errors.Is(err, certificates.ErrSpentOnGameNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err, certificateSpentOnGameNotFoundReason, certificateSpentOnGameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelCertificateToProtoCertificate(certificate), nil
}

func validateCreatedCertificate(certificate model.Certificate) error {
	return validation.ValidateStruct(&certificate,
		validation.Field(&certificate.Type, validation.Required, validation.By(model.ValidateCertificateType)),
		validation.Field(&certificate.WonOn, validation.Required, validation.Min(minWonOn)),
		validation.Field(&certificate.SpentOn, validation.By(validateSpentOn)),
		validation.Field(&certificate.Info, validation.By(validateCertificateInfo)),
	)
}
