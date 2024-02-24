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

// IsGameHasPhotos ...
func (i *Implementation) IsGameHasPhotos(ctx context.Context, req *photomanagerpb.IsGameHasPhotosRequest) (*photomanagerpb.IsGameHasPhotosResponse, error) {
	hasPhotos, err := i.gamePhotosFacade.IsGameHasPhotos(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &photomanagerpb.IsGameHasPhotosResponse{
		HasPhotos: hasPhotos,
	}, nil
}
