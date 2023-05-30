package registrator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	pkgmodel "github.com/nikita5637/quiz-registrator-api/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

// CreateCertificate ...
func (r *Registrator) CreateCertificate(ctx context.Context, req *registrator.CreateCertificateRequest) (*registrator.Certificate, error) {
	if err := validateCreateCertificateRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONInfoValue) {
			reason := fmt.Sprintf("invalid certificate info JSON value: \"%s\"", req.GetCertificate().GetInfo())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateInfoJSONValueLexeme)
		} else if errors.Is(err, errInvalidCertificateType) {
			reason := fmt.Sprintf("invalid certificate type: \"%d\"", req.GetCertificate().GetType())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateTypeLexeme)
		}

		return nil, st.Err()
	}

	certificate, err := r.certificatesFacade.CreateCertificate(ctx, model.Certificate{
		Type:    pkgmodel.CertificateType(req.GetCertificate().GetType()),
		WonOn:   req.GetCertificate().GetWonOn(),
		SpentOn: model.NewMaybeInt32(req.GetCertificate().GetSpentOn()),
		Info:    model.NewMaybeString(req.GetCertificate().GetInfo()),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrWonOnGameNotFound) {
			reason := fmt.Sprintf("won on game with id %d not found", req.GetCertificate().GetWonOn())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, wonOnGameNotFoundLexeme)
		} else if errors.Is(err, model.ErrSpentOnGameNotFound) {
			reason := fmt.Sprintf("spent on game with id %d not found", req.GetCertificate().GetSpentOn())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, spentOnGameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelCertificateToProtoCertificate(certificate), nil
}

// DeleteCertificate ...
func (r *Registrator) DeleteCertificate(ctx context.Context, req *registrator.DeleteCertificateRequest) (*emptypb.Empty, error) {
	err := r.certificatesFacade.DeleteCertificate(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrCertificateNotFound) {
			reason := fmt.Sprintf("certificate with ID %d not found", req.GetId())
			st = getStatus(ctx, codes.NotFound, err, reason, certificateNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}

// ListCertificates ...
func (r *Registrator) ListCertificates(ctx context.Context, _ *emptypb.Empty) (*registrator.ListCertificatesResponse, error) {
	certificates, err := r.certificatesFacade.ListCertificates(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respCertificates := make([]*registrator.Certificate, 0, len(certificates))
	for _, certificate := range certificates {
		respCertificates = append(respCertificates, &registrator.Certificate{
			Id:      certificate.ID,
			Type:    registrator.CertificateType(certificate.Type),
			WonOn:   certificate.WonOn,
			SpentOn: certificate.SpentOn.Value,
			Info:    certificate.Info.Value,
		})
	}

	return &registrator.ListCertificatesResponse{
		Certificates: respCertificates,
	}, nil
}

// PatchCertificate ...
func (r *Registrator) PatchCertificate(ctx context.Context, req *registrator.PatchCertificateRequest) (*registrator.Certificate, error) {
	if err := validatePatchCertificateRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONInfoValue) {
			reason := fmt.Sprintf("invalid certificate info JSON value: \"%s\"", req.GetCertificate().GetInfo())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateInfoJSONValueLexeme)
		} else if errors.Is(err, errInvalidCertificateType) {
			reason := fmt.Sprintf("invalid certificate type: \"%d\"", req.GetCertificate().GetType())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, invalidCertificateTypeLexeme)
		}

		return nil, st.Err()
	}

	certificate, err := r.certificatesFacade.PatchCertificate(ctx, model.Certificate{
		ID:      req.GetCertificate().GetId(),
		Type:    pkgmodel.CertificateType(req.GetCertificate().GetType()),
		WonOn:   req.GetCertificate().GetWonOn(),
		SpentOn: model.NewMaybeInt32(req.GetCertificate().GetSpentOn()),
		Info:    model.NewMaybeString(req.GetCertificate().GetInfo()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrCertificateNotFound) {
			reason := fmt.Sprintf("certificate with ID %d not found", req.GetCertificate().GetId())
			st = getStatus(ctx, codes.NotFound, err, reason, certificateNotFoundLexeme)
		} else if errors.Is(err, model.ErrWonOnGameNotFound) {
			reason := fmt.Sprintf("won on game with id %d not found", req.GetCertificate().GetWonOn())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, wonOnGameNotFoundLexeme)
		} else if errors.Is(err, model.ErrSpentOnGameNotFound) {
			reason := fmt.Sprintf("spent on game with id %d not found", req.GetCertificate().GetSpentOn())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, spentOnGameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelCertificateToProtoCertificate(certificate), nil
}

func convertModelCertificateToProtoCertificate(certificate model.Certificate) *registrator.Certificate {
	ret := &registrator.Certificate{
		Id:      certificate.ID,
		Type:    registrator.CertificateType(certificate.Type),
		WonOn:   certificate.WonOn,
		SpentOn: certificate.SpentOn.Value,
		Info:    certificate.Info.Value,
	}

	return ret
}

func validateCreateCertificateRequest(ctx context.Context, req *registrator.CreateCertificateRequest) error {
	return validateCertificate(ctx, req.GetCertificate())
}

func validatePatchCertificateRequest(ctx context.Context, req *registrator.PatchCertificateRequest) error {
	return validateCertificate(ctx, req.GetCertificate())
}

func validateCertificate(ctx context.Context, certificate *registrator.Certificate) error {
	if valid := json.Valid([]byte(certificate.GetInfo())); !valid {
		return errInvalidJSONInfoValue
	}

	err := validation.Validate(certificate.GetType(), validation.Required, validation.Min(1), validation.Max(int32(pkgmodel.NumberOfCertificateTypes-1)))
	if err != nil {
		return errInvalidCertificateType
	}

	if len(certificate.GetInfo()) > 256 {
		return errInvalidJSONInfoValue
	}

	return nil
}
