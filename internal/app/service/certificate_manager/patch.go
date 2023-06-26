package certificatemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PatchCertificate ...
func (m *CertificateManager) PatchCertificate(ctx context.Context, req *certificatemanagerpb.PatchCertificateRequest) (*certificatemanagerpb.Certificate, error) {
	if err := validatePatchCertificateRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONInfoValue) {
			reason := fmt.Sprintf("invalid certificate info JSON value: \"%s\"", req.GetCertificate().GetInfo())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateInfoJSONValueLexeme)
		} else if errors.Is(err, errInvalidCertificateType) {
			reason := fmt.Sprintf("invalid certificate type: \"%d\"", req.GetCertificate().GetType())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateTypeLexeme)
		}

		return nil, st.Err()
	}

	certificate, err := m.certificatesFacade.PatchCertificate(ctx, model.Certificate{
		ID:      req.GetCertificate().GetId(),
		Type:    model.CertificateType(req.GetCertificate().GetType()),
		WonOn:   req.GetCertificate().GetWonOn(),
		SpentOn: model.NewMaybeInt32(req.GetCertificate().GetSpentOn()),
		Info:    model.NewMaybeString(req.GetCertificate().GetInfo()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrCertificateNotFound) {
			reason := fmt.Sprintf("certificate with ID %d not found", req.GetCertificate().GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, certificateNotFoundLexeme)
		} else if errors.Is(err, model.ErrWonOnGameNotFound) {
			reason := fmt.Sprintf("won on game with id %d not found", req.GetCertificate().GetWonOn())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, wonOnGameNotFoundLexeme)
		} else if errors.Is(err, model.ErrSpentOnGameNotFound) {
			reason := fmt.Sprintf("spent on game with id %d not found", req.GetCertificate().GetSpentOn())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, spentOnGameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelCertificateToProtoCertificate(certificate), nil
}

func validatePatchCertificateRequest(ctx context.Context, req *certificatemanagerpb.PatchCertificateRequest) error {
	return validateCertificate(ctx, req.GetCertificate())
}
