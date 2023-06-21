package certificatemanager

import (
	"context"

	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListCertificates ...
func (m *CertificateManager) ListCertificates(ctx context.Context, _ *emptypb.Empty) (*certificatemanagerpb.ListCertificatesResponse, error) {
	certificates, err := m.certificatesFacade.ListCertificates(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respCertificates := make([]*certificatemanagerpb.Certificate, 0, len(certificates))
	for _, certificate := range certificates {
		respCertificates = append(respCertificates, &certificatemanagerpb.Certificate{
			Id:      certificate.ID,
			Type:    certificatemanagerpb.CertificateType(certificate.Type),
			WonOn:   certificate.WonOn,
			SpentOn: certificate.SpentOn.Value,
			Info:    certificate.Info.Value,
		})
	}

	return &certificatemanagerpb.ListCertificatesResponse{
		Certificates: respCertificates,
	}, nil
}
