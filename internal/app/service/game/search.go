package game

import (
	"context"

	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SearchGamesByLeagueID ...
func (i *Implementation) SearchGamesByLeagueID(ctx context.Context, req *gamepb.SearchGamesByLeagueIDRequest) (*gamepb.SearchGamesByLeagueIDResponse, error) {
	page := req.GetPage()
	if page > 0 {
		page--
	}

	offset := page * req.GetPageSize()

	games, total, err := i.gamesFacade.SearchGamesByLeagueID(ctx, int32(req.GetId()), offset, req.GetPageSize())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*gamepb.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToProtoGame(game))
	}

	return &gamepb.SearchGamesByLeagueIDResponse{
		Games: pbGames,
		Total: total,
	}, nil
}

// SearchPassedAndRegisteredGames ...
func (i *Implementation) SearchPassedAndRegisteredGames(ctx context.Context, req *gamepb.SearchPassedAndRegisteredGamesRequest) (*gamepb.SearchPassedAndRegisteredGamesResponse, error) {
	page := req.GetPage()
	if page > 0 {
		page--
	}

	offset := page * req.GetPageSize()

	games, total, err := i.gamesFacade.SearchPassedAndRegisteredGames(ctx, offset, req.GetPageSize())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*gamepb.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToProtoGame(game))
	}

	return &gamepb.SearchPassedAndRegisteredGamesResponse{
		Games: pbGames,
		Total: total,
	}, nil
}
