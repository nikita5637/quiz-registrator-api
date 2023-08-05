package gameplayer

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UnregisterPlayer ...
func (i *Implementation) UnregisterPlayer(ctx context.Context, req *gameplayerpb.UnregisterPlayerRequest) (*emptypb.Empty, error) {
	if req.GetGamePlayer() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	unregisteredGamePlayer := convertProtoGamePlayerToModelGamePlayer(req.GetGamePlayer())

	logger.Debugf(ctx, "trying to unregister game player: %#v", unregisteredGamePlayer)

	if err := validateUnregisteredGamePlayer(unregisteredGamePlayer); err != nil {
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

	existedGamePlayers, err := i.gamePlayersFacade.GetGamePlayersByFields(ctx, unregisteredGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if len(existedGamePlayers) == 0 {
		st := status.New(codes.NotFound, gameplayers.ErrGamePlayerNotFound.Error())
		return nil, st.Err()
	}

	_, err = i.gamesFacade.GetGameByID(ctx, req.GetGamePlayer().GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx,
				codes.FailedPrecondition,
				games.ErrGameNotFound.Error(),
				games.ReasonGameNotFound,
				map[string]string{
					"error": err.Error(),
				},
				games.GameNotFoundLexeme,
			)
		} else if errors.Is(err, games.ErrGameHasPassed) {
			st = model.GetStatus(ctx,
				codes.FailedPrecondition,
				games.ErrGameHasPassed.Error(),
				games.ReasonGameHasPassed,
				map[string]string{
					"error": err.Error(),
				},
				games.GameHasPassedLexeme,
			)

		}

		return nil, st.Err()
	}

	err = i.gamePlayersFacade.DeleteGamePlayer(ctx, existedGamePlayers[0].ID)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerNotFound) {
			st = model.GetStatus(ctx,
				codes.NotFound,
				gameplayers.ErrGamePlayerNotFound.Error(),
				gameplayers.ReasonGamePlayerNotFound,
				nil,
				gameplayers.GamePlayerNotFoundLexeme,
			)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}

func validateUnregisteredGamePlayer(gamePlayer model.GamePlayer) error {
	return validateCreatedGamePlayer(gamePlayer)
}
