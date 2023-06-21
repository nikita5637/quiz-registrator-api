package certificatemanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteCertificate ...
func (m *CertificateManager) DeleteCertificate(ctx context.Context, req *certificatemanagerpb.DeleteCertificateRequest) (*emptypb.Empty, error) {
	err := m.certificatesFacade.DeleteCertificate(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrCertificateNotFound) {
			reason := fmt.Sprintf("certificate with ID %d not found", req.GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, certificateNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
