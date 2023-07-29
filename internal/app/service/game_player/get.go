package gameplayer

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetGamePlayer ...
func (i *Implementation) GetGamePlayer(ctx context.Context, req *gameplayerpb.GetGamePlayerRequest) (*gameplayerpb.GamePlayer, error) {
	gamePlayer, err := i.gamePlayersFacade.GetGamePlayer(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, gameplayers.ErrGamePlayerNotFound, gameplayers.ReasonGamePlayerNotFound, gameplayers.GamePlayerNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGamePlayerToProtoGamePlayer(gamePlayer), nil
}

// GetGamePlayersByGameID ...
func (i *Implementation) GetGamePlayersByGameID(ctx context.Context, req *gameplayerpb.GetGamePlayersByGameIDRequest) (*gameplayerpb.GetGamePlayersByGameIDResponse, error) {
	gamePlayers, err := i.gamePlayersFacade.GetGamePlayersByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGamePlayers := make([]*gameplayerpb.GamePlayer, 0, len(gamePlayers))
	for _, gamePlayer := range gamePlayers {
		pbGamePlayers = append(pbGamePlayers, convertModelGamePlayerToProtoGamePlayer(gamePlayer))
	}

	return &gameplayerpb.GetGamePlayersByGameIDResponse{
		GamePlayers: pbGamePlayers,
	}, nil
}
