package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdatePayment ...
func (r *Registrator) UpdatePayment(ctx context.Context, req *registrator.UpdatePaymentRequest) (*registrator.UpdatePaymentResponse, error) {
	err := r.gamesFacade.UpdatePayment(ctx, req.GetGameId(), int32(req.GetPayment()))
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.UpdatePaymentResponse{}, nil
}
