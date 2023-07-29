package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterGame ...
func (r *Registrator) RegisterGame(ctx context.Context, req *registrator.RegisterGameRequest) (*registrator.RegisterGameResponse, error) {
	registerStatus, err := r.gamesFacade.RegisterGame(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, games.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.RegisterGameResponse{
		Status: registrator.RegisterGameStatus(registerStatus),
	}, nil
}
