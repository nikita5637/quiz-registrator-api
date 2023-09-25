package game

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BatchGetGames ...
func (i *Implementation) BatchGetGames(ctx context.Context, req *gamepb.BatchGetGamesRequest) (*gamepb.BatchGetGamesResponse, error) {
	games, err := i.gamesFacade.GetGamesByIDs(ctx, req.GetIds())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*gamepb.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToProtoGame(game))
	}
	return &gamepb.BatchGetGamesResponse{
		Games: pbGames,
	}, nil
}

// GetGame ...
func (i *Implementation) GetGame(ctx context.Context, req *gamepb.GetGameRequest) (*gamepb.Game, error) {
	game, err := i.gamesFacade.GetGameByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameToProtoGame(game), nil
}
