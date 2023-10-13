package gameplayer

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteGamePlayer ...
func (i *Implementation) DeleteGamePlayer(ctx context.Context, req *gameplayerpb.DeleteGamePlayerRequest) (*emptypb.Empty, error) {
	err := i.gamePlayersFacade.DeleteGamePlayer(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, gameplayers.ErrGamePlayerNotFound.Error(), gameplayers.ReasonGamePlayerNotFound, map[string]string{
				"error": err.Error(),
			}, gameplayers.GamePlayerNotFoundLexeme)
		}

		return &emptypb.Empty{}, st.Err()
	}

	return &emptypb.Empty{}, nil
}
