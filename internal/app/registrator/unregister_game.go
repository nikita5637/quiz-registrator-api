package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnregisterGame ...
func (r *Registrator) UnregisterGame(ctx context.Context, req *registrator.UnregisterGameRequest) (*registrator.UnregisterGameResponse, error) {
	unregisterStatus, err := r.gamesFacade.UnregisterGame(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.UnregisterGameResponse{
		Status: registrator.UnregisterGameStatus(unregisterStatus),
	}, nil
}
