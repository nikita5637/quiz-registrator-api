package registrator

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	noFreeSlotsLexeme = i18n.Lexeme{
		Key:      "no_free_slots",
		FallBack: "There are not free slots",
	}
	unauthenticatedRequestLexeme = i18n.Lexeme{
		Key:      "unauthenticated_request",
		FallBack: "Unauthenticated request",
	}
)

// RegisterPlayer ...
func (r *Registrator) RegisterPlayer(ctx context.Context, req *registrator.RegisterPlayerRequest) (*registrator.RegisterPlayerResponse, error) {
	user := users_utils.UserFromContext(ctx)
	if user.ID == 0 {
		err := errors.New("unauthenticated")
		st := getStatus(ctx, codes.Unauthenticated, err, unauthenticatedRequestReason, unauthenticatedRequestLexeme)
		return nil, st.Err()
	}

	userID := user.ID
	registeredBy := user.ID
	if req.GetPlayerType() == registrator.PlayerType_PLAYER_TYPE_LEGIONER {
		userID = 0
	}

	registerStatus, err := r.gamesFacade.RegisterPlayer(ctx, req.GetGameId(), userID, registeredBy, int32(req.GetDegree()))
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		} else if errors.Is(err, model.ErrGameNoFreeSlots) {
			reason := fmt.Sprintf("no free slots for game with id %d", req.GetGameId())
			st = getStatus(ctx, codes.AlreadyExists, err, reason, noFreeSlotsLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.RegisterPlayerResponse{
		Status: registrator.RegisterPlayerStatus(registerStatus),
	}, nil
}
