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

// GetPhotosByGameID ...
func (i *Implementation) GetPhotosByGameID(ctx context.Context, req *photomanagerpb.GetPhotosByGameIDRequest) (*photomanagerpb.GetPhotosByGameIDResponse, error) {
	urls, err := i.gamePhotosFacade.GetPhotosByGameID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &photomanagerpb.GetPhotosByGameIDResponse{
		Urls: urls,
	}, nil
}
