package registrator

import (
	"errors"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRegistrator_UpdatePayment(t *testing.T) {
	t.Run("internal error while get game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate).Return(errors.New("some error"))

		got, err := fx.registrator.UpdatePayment(fx.ctx, &registrator.UpdatePaymentRequest{
			GameId:  1,
			Payment: commonpb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate).Return(games.ErrGameNotFound)

		got, err := fx.registrator.UpdatePayment(fx.ctx, &registrator.UpdatePaymentRequest{
			GameId:  1,
			Payment: commonpb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().UpdatePayment(fx.ctx, int32(1), model.PaymentCertificate).Return(nil)

		got, err := fx.registrator.UpdatePayment(fx.ctx, &registrator.UpdatePaymentRequest{
			GameId:  1,
			Payment: commonpb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.NotNil(t, got)
		assert.Equal(t, &registrator.UpdatePaymentResponse{}, got)
		assert.NoError(t, err)
	})
}
