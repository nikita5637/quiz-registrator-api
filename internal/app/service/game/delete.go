package game

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteGame ...
func (i *Implementation) DeleteGame(ctx context.Context, req *gamepb.DeleteGameRequest) (*emptypb.Empty, error) {
	if err := i.gamesFacade.DeleteGame(ctx, req.GetId()); err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
