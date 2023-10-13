package gameplayer

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateGamePlayer ...
func (i *Implementation) CreateGamePlayer(ctx context.Context, req *gameplayerpb.CreateGamePlayerRequest) (*gameplayerpb.GamePlayer, error) {
	if req.GetGamePlayer() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	createdGamePlayer := convertProtoGamePlayerToModelGamePlayer(req.GetGamePlayer())

	logger.DebugKV(ctx, "creating new game player", "gameplayer", createdGamePlayer)

	if err := validateCreatedGamePlayer(createdGamePlayer); err != nil {
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

	gamePlayer, err := i.gamePlayersFacade.CreateGamePlayer(ctx, createdGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, gameplayers.ErrGamePlayerAlreadyExists.Error(), gameplayers.ReasonGamePlayerAlreadyExists, map[string]string{
				"error": err.Error(),
			}, gameplayers.GamePlayerAlreadyExistsLexeme)
		} else if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		} else if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, users.ErrUserNotFound.Error(), users.ReasonUserNotFound, map[string]string{
				"error": err.Error(),
			}, users.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGamePlayerToProtoGamePlayer(gamePlayer), nil
}

func validateCreatedGamePlayer(gamePlayer model.GamePlayer) error {
	return validation.ValidateStruct(&gamePlayer,
		validation.Field(&gamePlayer.GameID, validation.Required, validation.Min(1)),
		validation.Field(&gamePlayer.UserID, validation.By(validateUserID)),
		validation.Field(&gamePlayer.RegisteredBy, validation.Required, validation.Min(1)),
		validation.Field(&gamePlayer.Degree, validation.By(model.ValidateDegree)),
	)
}
