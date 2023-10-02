package certificatemanager

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/certificates"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	minID = int32(1)
)

// PatchCertificate ...
func (m *CertificateManager) PatchCertificate(ctx context.Context, req *certificatemanagerpb.PatchCertificateRequest) (*certificatemanagerpb.Certificate, error) {
	originalCertificate, err := m.certificatesFacade.GetCertificate(ctx, req.GetCertificate().GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, certificates.ErrCertificateNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, err.Error(), "CERTIFICATE_NOT_FOUND", nil, certificateNotFoundLexeme)
		}

		return nil, st.Err()
	}

	patchedCertificate := originalCertificate
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "type":
			patchedCertificate.Type = model.CertificateType(req.GetCertificate().GetType())
		case "won_on":
			patchedCertificate.WonOn = req.GetCertificate().GetWonOn()
		case "spent_on":
			if req.GetCertificate().GetSpentOn() != nil {
				patchedCertificate.SpentOn = maybe.Just(req.GetCertificate().GetSpentOn().GetValue())
			}
		case "info":
			if req.GetCertificate().GetInfo() != nil {
				patchedCertificate.Info = maybe.Just(req.GetCertificate().GetInfo().GetValue())
			}
		}
	}

	if err = validatePatchedCertificate(patchedCertificate); err != nil {
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

	certificate, err := m.certificatesFacade.PatchCertificate(ctx, patchedCertificate)
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

func validatePatchedCertificate(certificate model.Certificate) error {
	return validation.ValidateStruct(&certificate,
		validation.Field(&certificate.ID, validation.Required, validation.Min(minID)),
		validation.Field(&certificate.Type, validation.Required, validation.By(model.ValidateCertificateType)),
		validation.Field(&certificate.WonOn, validation.Required, validation.Min(minWonOn)),
		validation.Field(&certificate.SpentOn, validation.By(validateSpentOn)),
		validation.Field(&certificate.Info, validation.By(validateCertificateInfo)),
	)
}
