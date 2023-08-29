package photomanager

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetGameWithPhotosIDs ...
func (i *Implementation) GetGameWithPhotosIDs(ctx context.Context, req *photomanagerpb.GetGameWithPhotosIDsRequest) (*photomanagerpb.GetGameWithPhotosIDsResponse, error) {
	gamesTotal, err := i.gamePhotosFacade.GetNumberOfGamesWithPhotos(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if gamesTotal == 0 {
		return &photomanagerpb.GetGameWithPhotosIDsResponse{}, nil
	}

	gameIDs, err := i.gamePhotosFacade.GetGameWithPhotosIDs(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	return &photomanagerpb.GetGameWithPhotosIDsResponse{
		GameIds: gameIDs,
		Total:   gamesTotal,
	}, nil
}

// GetPhotosByGameID ...
func (i *Implementation) GetPhotosByGameID(ctx context.Context, req *photomanagerpb.GetPhotosByGameIDRequest) (*photomanagerpb.GetPhotosByGameIDResponse, error) {
	urls, err := i.gamePhotosFacade.GetPhotosByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, err.Error(), games.ReasonGameNotFound, nil, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &photomanagerpb.GetPhotosByGameIDResponse{
		Urls: urls,
	}, nil
}
