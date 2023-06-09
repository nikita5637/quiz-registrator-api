package certificatemanager

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	certificatemanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/certificate_manager"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRegistrator_DeleteCertificate(t *testing.T) {
	t.Run("error certificate not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(model.ErrCertificateNotFound)

		got, err := fx.certificateManager.DeleteCertificate(fx.ctx, &certificatemanagerpb.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("some internal error", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(errors.New("some error"))

		got, err := fx.certificateManager.DeleteCertificate(fx.ctx, &certificatemanagerpb.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.certificatesFacade.EXPECT().DeleteCertificate(fx.ctx, int32(1)).Return(nil)

		got, err := fx.certificateManager.DeleteCertificate(fx.ctx, &certificatemanagerpb.DeleteCertificateRequest{
			Id: 1,
		})

		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
