package photomanager

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	photomanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/photo_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AddGamePhotos ...
func (i *Implementation) AddGamePhotos(ctx context.Context, req *photomanagerpb.AddGamePhotosRequest) (*emptypb.Empty, error) {
	err := i.gamePhotosFacade.AddGamePhotos(ctx, req.GetGameId(), req.GetUrls())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, games.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
