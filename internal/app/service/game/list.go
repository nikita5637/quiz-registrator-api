package game

import (
	"context"

	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListGames ...
func (i *Implementation) ListGames(ctx context.Context, req *emptypb.Empty) (*gamepb.ListGamesResponse, error) {
	games, err := i.gamesFacade.ListGames(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*gamepb.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToProtoGame(game))
	}

	return &gamepb.ListGamesResponse{
		Games: pbGames,
	}, nil
}
