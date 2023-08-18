package gameplayer

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameplayers"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/users"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	reasonThereAreNoFreeSlot = "THERE_ARE_NO_FREE_SLOT"
)

var (
	noFreeSlotLexeme = i18n.Lexeme{
		Key:      "no_free_slot",
		FallBack: "There are not free slot",
	}
)

// RegisterPlayer ...
func (i *Implementation) RegisterPlayer(ctx context.Context, req *gameplayerpb.RegisterPlayerRequest) (*emptypb.Empty, error) {
	if req.GetGamePlayer() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	registeredGamePlayer := convertProtoGamePlayerToModelGamePlayer(req.GetGamePlayer())

	logger.Debugf(ctx, "trying to register new game player: %#v", registeredGamePlayer)

	if err := validateRegisteredGamePlayer(registeredGamePlayer); err != nil {
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

	existedGamePlayers, err := i.gamePlayersFacade.GetGamePlayersByGameID(ctx, req.GetGamePlayer().GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	game, err := i.gamesFacade.GetGameByID(ctx, req.GetGamePlayer().GetGameId())
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

	if len(existedGamePlayers) >= int(game.MaxPlayers) {
		st := model.GetStatus(ctx, codes.FailedPrecondition, "there are no free slot", reasonThereAreNoFreeSlot, nil, noFreeSlotLexeme)
		return nil, st.Err()
	}

	_, err = i.gamePlayersFacade.CreateGamePlayer(ctx, registeredGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerAlreadyRegistered) {
			st = model.GetStatus(ctx, codes.AlreadyExists, gameplayers.ErrGamePlayerAlreadyRegistered.Error(), gameplayers.ReasonGamePlayerAlreadyRegistered, nil, gameplayers.GamePlayerAlreadyRegisteredLexeme)
		} else if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx,
				codes.FailedPrecondition,
				games.ErrGameNotFound.Error(),
				games.ReasonGameNotFound,
				map[string]string{
					"error": err.Error(),
				},
				games.GameNotFoundLexeme,
			)

		} else if errors.Is(err, users.ErrUserNotFound) {
			st = model.GetStatus(ctx,
				codes.FailedPrecondition,
				users.ErrUserNotFound.Error(),
				users.ReasonUserNotFound,
				map[string]string{
					"error": err.Error(),
				},
				users.UserNotFoundLexeme,
			)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}

func validateRegisteredGamePlayer(gamePlayer model.GamePlayer) error {
	return validateCreatedGamePlayer(gamePlayer)
}