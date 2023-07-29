package croupier

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
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
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		} else if errors.Is(err, games.ErrGameHasPassed) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	if !game.My {
		return &croupierpb.GetLotteryStatusResponse{
			Active: false,
		}, nil
	}

	return &croupierpb.GetLotteryStatusResponse{
		Active: i.croupier.GetIsLotteryActive(ctx, game),
	}, nil
}
