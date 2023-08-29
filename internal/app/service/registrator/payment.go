package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdatePayment ...
func (i *Implementation) UpdatePayment(ctx context.Context, req *registrator.UpdatePaymentRequest) (*registrator.UpdatePaymentResponse, error) {
	err := i.gamesFacade.UpdatePayment(ctx, req.GetGameId(), model.PaymentType(req.GetPayment()))
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, games.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.UpdatePaymentResponse{}, nil
}
