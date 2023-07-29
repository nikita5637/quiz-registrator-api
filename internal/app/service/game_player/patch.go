package gameplayer

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	minID           = int32(1)
	minGameID       = int32(1)
	minRegisteredBy = int32(1)
	minUserID       = int32(1)
)

// PatchGamePlayer ...
func (i *Implementation) PatchGamePlayer(ctx context.Context, req *gameplayerpb.PatchGamePlayerRequest) (*gameplayerpb.GamePlayer, error) {
	if req.GetGamePlayer() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	originalGamePlayer, err := i.gamePlayersFacade.GetGamePlayer(ctx, req.GetGamePlayer().GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, err.Error(), gameplayers.ReasonGamePlayerNotFound, nil, gameplayers.GamePlayerNotFoundLexeme)
		}

		return nil, st.Err()
	}

	patchedGamePlayer := originalGamePlayer
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "game_id":
			patchedGamePlayer.GameID = req.GetGamePlayer().GetGameId()
		case "user_id":
			if req.GetGamePlayer().GetUserId() != nil {
				patchedGamePlayer.UserID = maybe.Just(req.GetGamePlayer().GetUserId().GetValue())
			}
		case "registered_by":
			patchedGamePlayer.RegisteredBy = req.GetGamePlayer().GetRegisteredBy()
		case "degree":
			patchedGamePlayer.Degree = model.Degree(req.GetGamePlayer().GetDegree())
		}
	}

	if err = validatePatchedGamePlayer(patchedGamePlayer); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if validationErrors, ok := err.(validation.Errors); ok && len(validationErrors) > 0 {
			keys := make([]string, 0, len(validationErrors))
			for k := range validationErrors {
				keys = append(keys, k)
			}

			if errorDetails := getErrorDetails(keys); errorDetails != nil {
				st = model.GetStatus(ctx,
					codes.InvalidArgument,
					fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()),
					errorDetails.Reason,
					map[string]string{
						"error": err.Error(),
					},
					errorDetails.Lexeme,
				)
			}
		}

		return nil, st.Err()
	}

	gamePlayer, err := i.gamePlayersFacade.PatchGamePlayer(ctx, patchedGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), users.ReasonUserNotFound, nil, users.UserNotFoundLexeme)
		} else if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), games.ReasonGameNotFound, nil, games.GameNotFoundLexeme)
		} else if errors.Is(err, gameplayers.ErrGamePlayerAlreadyRegistered) {
			st = model.GetStatus(ctx, codes.AlreadyExists, err.Error(), gameplayers.ReasonGamePlayerAlreadyRegistered, nil, gameplayers.GamePlayerAlreadyRegisteredLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGamePlayerToProtoGamePlayer(gamePlayer), nil
}

func validatePatchedGamePlayer(gamePlayer model.GamePlayer) error {
	return validation.ValidateStruct(&gamePlayer,
		validation.Field(&gamePlayer.ID, validation.Required, validation.Min(minID)),
		validation.Field(&gamePlayer.GameID, validation.Required, validation.Min(minGameID)),
		validation.Field(&gamePlayer.UserID, validation.By(validateUserID)),
		validation.Field(&gamePlayer.RegisteredBy, validation.Required, validation.Min(minRegisteredBy)),
		validation.Field(&gamePlayer.Degree, validation.By(model.ValidateDegree)),
	)
}
