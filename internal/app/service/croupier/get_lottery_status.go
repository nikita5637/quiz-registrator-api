package croupier

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetLotteryStatus ...
func (i *Implemintation) GetLotteryStatus(ctx context.Context, req *croupierpb.GetLotteryStatusRequest) (*croupierpb.GetLotteryStatusResponse, error) {
	game, err := i.gamesFacade.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	if game.HasPassed {
		return &croupierpb.GetLotteryStatusResponse{
			Active: false,
		}, nil
	}

	return &croupierpb.GetLotteryStatusResponse{
		Active: i.croupier.GetIsLotteryActive(ctx, game),
	}, nil
}
