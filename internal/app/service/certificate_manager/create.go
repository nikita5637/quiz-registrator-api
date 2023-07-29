package certificatemanager

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCertificate ...
func (m *CertificateManager) CreateCertificate(ctx context.Context, req *certificatemanagerpb.CreateCertificateRequest) (*certificatemanagerpb.Certificate, error) {
	createdCertificate := convertProtoCertificateToModelCertificate(req.GetCertificate())
	if err := validateCreatedCertificate(createdCertificate); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if validationErrors, ok := err.(validation.Errors); ok && len(validationErrors) > 0 {
			keys := make([]string, 0, len(validationErrors))
			for k := range validationErrors {
				keys = append(keys, k)
			}

			if errorDetails := getErrorDetails(keys); errorDetails != nil {
				st = model.GetStatus(ctx,
					codes.InvalidArgument,
					fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()),
					errorDetails.Reason,
					map[string]string{
						"error": err.Error(),
					},
					errorDetails.Lexeme,
				)
			}
		}

		return nil, st.Err()
	}

	certificate, err := m.certificatesFacade.CreateCertificate(ctx, createdCertificate)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, certificates.ErrWonOnGameNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), reasonCertificateWonOnGameNotFound, nil, certificateWonOnGameNotFoundLexeme)
		} else if errors.Is(err, certificates.ErrSpentOnGameNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), reasonCertificateSpentOnGameNotFound, nil, certificateSpentOnGameNotFoundLexeme)
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
