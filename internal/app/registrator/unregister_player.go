package registrator

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnregisterPlayer ...
func (r *Registrator) UnregisterPlayer(ctx context.Context, req *registrator.UnregisterPlayerRequest) (*registrator.UnregisterPlayerResponse, error) {
	user := users_utils.UserFromContext(ctx)
	if user.ID == 0 {
		err := errors.New("unauthenticated")
		st := getStatus(ctx, codes.Unauthenticated, err, unauthenticatedRequestReason, unauthenticatedRequestLexeme)
		return nil, st.Err()
	}

	userID := user.ID
	if req.GetPlayerType() == registrator.PlayerType_PLAYER_TYPE_LEGIONER {
		userID = 0
	}

	unregisterStatus, err := r.gamesFacade.UnregisterPlayer(ctx, req.GetGameId(), userID, user.ID)
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.UnregisterPlayerResponse{
		Status: registrator.UnregisterPlayerStatus(unregisterStatus),
	}, nil
}
