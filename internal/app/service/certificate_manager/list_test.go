package certificatemanager

import (
	"errors"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestRegistrator_ListCertificates(t *testing.T) {
	t.Run("error while list certificates", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().ListCertificates(fx.ctx).Return(nil, errors.New("some error"))

		got, err := fx.certificateManager.ListCertificates(fx.ctx, &emptypb.Empty{})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().ListCertificates(fx.ctx).Return([]model.Certificate{
			{
				ID:      1,
				Type:    model.CertificateTypeFreePass,
				WonOn:   1,
				SpentOn: maybe.Just(int32(2)),
				Info:    maybe.Just("{}"),
			},
			{
				ID:      2,
				Type:    model.CertificateTypeBarBillPayment,
				WonOn:   3,
				SpentOn: maybe.Just(int32(2)),
				Info:    maybe.Just("{}"),
			},
			{
				ID:      3,
				Type:    model.CertificateTypeBarBillPayment,
				WonOn:   4,
				SpentOn: maybe.Nothing[int32](),
				Info:    maybe.Nothing[string](),
			},
		}, nil)

		got, err := fx.certificateManager.ListCertificates(fx.ctx, &emptypb.Empty{})
		assert.ElementsMatch(t,
			[]*certificatemanagerpb.Certificate{
				{
					Id:      1,
					Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_FREE_PASS,
					WonOn:   1,
					SpentOn: wrapperspb.Int32(2),
					Info:    wrapperspb.String("{}"),
				},
				{
					Id:      2,
					Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
					WonOn:   3,
					SpentOn: wrapperspb.Int32(2),
					Info:    wrapperspb.String("{}"),
				},
				{
					Id:      3,
					Type:    certificatemanagerpb.CertificateType_CERTIFICATE_TYPE_BAR_BILL_PAYMENT,
					WonOn:   4,
					SpentOn: nil,
					Info:    nil,
				},
			},
			got.GetCertificates())
		assert.NoError(t, err)
	})
}
