package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddGamePhotos ...
func (r *Registrator) AddGamePhotos(ctx context.Context, req *registrator.AddGamePhotosRequest) (*registrator.AddGamePhotosResponse, error) {
	err := r.gamePhotosFacade.AddGamePhotos(ctx, req.GetGameId(), req.GetUrls())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.AddGamePhotosResponse{}, nil
}

// GetGamesWithPhotos ...
func (r *Registrator) GetGamesWithPhotos(ctx context.Context, req *registrator.GetGamesWithPhotosRequest) (*registrator.GetGamesWithPhotosResponse, error) {
	gamesTotal, err := r.gamePhotosFacade.GetNumberOfGamesWithPhotos(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if gamesTotal == 0 {
		return &registrator.GetGamesWithPhotosResponse{}, nil
	}

	games, err := r.gamePhotosFacade.GetGamesWithPhotos(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	pbGames := make([]*registrator.Game, 0, len(games))
	for _, game := range games {
		pbGames = append(pbGames, convertModelGameToPBGame(game))
	}

	return &registrator.GetGamesWithPhotosResponse{
		Games: pbGames,
		Total: gamesTotal,
	}, nil
}

// GetPhotosByGameID ...
func (r *Registrator) GetPhotosByGameID(ctx context.Context, req *registrator.GetPhotosByGameIDRequest) (*registrator.GetPhotosByGameIDResponse, error) {
	urls, err := r.gamePhotosFacade.GetPhotosByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.GetPhotosByGameIDResponse{
		Urls: urls,
	}, nil
}
