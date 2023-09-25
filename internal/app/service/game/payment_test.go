package game

import (
	"errors"
	"math"
	"testing"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestImplementation_UpdatePayment(t *testing.T) {
	t.Run("error: game not found", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, games.ErrGameNotFound)

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.NotFound, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_NOT_FOUND", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "game not found",
		}, errorInfo.Metadata)
	})

	t.Run("error: internal error while getting game", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("error: game has passed", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: true,
		}, nil)

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.FailedPrecondition, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "GAME_HAS_PASSED", errorInfo.Reason)
		assert.Nil(t, errorInfo.Metadata)
	})

	t.Run("error: validation error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment(math.MaxInt32),
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.InvalidArgument, st.Code())
		assert.Len(t, st.Details(), 2)

		errorInfo, ok := st.Details()[0].(*errdetails.ErrorInfo)
		assert.True(t, ok)
		assert.Equal(t, "INVALID_PAYMENT", errorInfo.Reason)
		assert.Equal(t, map[string]string{
			"error": "Payment: must be no greater than 3.",
		}, errorInfo.Metadata)
	})

	t.Run("error: patch game error", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:        1,
			Payment:   maybe.Nothing[model.Payment](),
			HasPassed: false,
		}).Return(model.Game{}, errors.New("some error"))

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment_PAYMENT_INVALID,
		})
		assert.Nil(t, got)
		assert.Error(t, err)

		st := status.Convert(err)
		assert.Equal(t, codes.Internal, st.Code())
		assert.Len(t, st.Details(), 0)
	})

	t.Run("ok", func(t *testing.T) {
		fx := tearUp(t)

		fx.gamesFacade.EXPECT().GetGameByID(fx.ctx, int32(1)).Return(model.Game{
			ID:        1,
			HasPassed: false,
		}, nil)

		fx.gamesFacade.EXPECT().PatchGame(fx.ctx, model.Game{
			ID:        1,
			Payment:   maybe.Just(model.PaymentCertificate),
			HasPassed: false,
		}).Return(model.Game{
			ID:        1,
			Payment:   maybe.Just(model.PaymentCertificate),
			HasPassed: false,
		}, nil)

		got, err := fx.implementation.UpdatePayment(fx.ctx, &gamepb.UpdatePaymentRequest{
			Id:      1,
			Payment: gamepb.Payment_PAYMENT_CERTIFICATE,
		})
		assert.Equal(t, &emptypb.Empty{}, got)
		assert.NoError(t, err)
	})
}
