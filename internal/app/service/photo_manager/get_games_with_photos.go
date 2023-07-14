package photomanager

import (
	"context"

	commonpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/common"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetGamesWithPhotos ...
func (i *Implementation) GetGamesWithPhotos(ctx context.Context, req *photomanagerpb.GetGamesWithPhotosRequest) (*photomanagerpb.GetGamesWithPhotosResponse, error) {
	gamesTotal, err := i.gamePhotosFacade.GetNumberOfGamesWithPhotos(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if gamesTotal == 0 {
		return &photomanagerpb.GetGamesWithPhotosResponse{}, nil
	}

	games, err := i.gamePhotosFacade.GetGamesWithPhotos(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*commonpb.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToPBGame(game))
	}

	return &photomanagerpb.GetGamesWithPhotosResponse{
		Games: pbGames,
		Total: gamesTotal,
	}, nil
}
