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
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	reasonThereAreNoFreeSlots = "THERE_ARE_NO_FREE_SLOTS"
)

var (
	noFreeSlotsLexeme = i18n.Lexeme{
		Key:      "no_free_slots",
		FallBack: "There are not free slots",
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
				st = status.New(codes.InvalidArgument, fmt.Sprintf("%s %s", keys[0], validationErrors[keys[0]].Error()))
				errorInfo := &errdetails.ErrorInfo{
					Reason: errorDetails.Reason,
					Metadata: map[string]string{
						"error": err.Error(),
					},
				}
				localizedMessage := &errdetails.LocalizedMessage{
					Locale:  i18n.GetLangFromContext(ctx),
					Message: i18n.GetTranslator(errorDetails.Lexeme)(ctx),
				}
				st, _ = st.WithDetails(errorInfo, localizedMessage)
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
			st = status.New(codes.InvalidArgument, games.ErrGameNotFound.Error())
			errorInfo := &errdetails.ErrorInfo{
				Reason: games.ReasonGameNotFound,
				Metadata: map[string]string{
					"error": err.Error(),
				},
			}
			localizedMessage := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(games.GameNotFoundLexeme)(ctx),
			}
			st, _ = st.WithDetails(errorInfo, localizedMessage)
		}

		return nil, st.Err()
	}

	if len(existedGamePlayers) >= int(game.MaxPlayers) {
		st := status.New(codes.FailedPrecondition, "there are no free slots")
		errorInfo := &errdetails.ErrorInfo{
			Reason: reasonThereAreNoFreeSlots,
		}
		localizedMessage := &errdetails.LocalizedMessage{
			Locale:  i18n.GetLangFromContext(ctx),
			Message: i18n.GetTranslator(noFreeSlotsLexeme)(ctx),
		}
		st, _ = st.WithDetails(errorInfo, localizedMessage)

		return nil, st.Err()
	}

	_, err = i.gamePlayersFacade.CreateGamePlayer(ctx, registeredGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerAlreadyRegistered) {
			st = status.New(codes.AlreadyExists, gameplayers.ErrGamePlayerAlreadyRegistered.Error())
			errorInfo := &errdetails.ErrorInfo{
				Reason: gameplayers.ReasonGamePlayerAlreadyRegistered,
			}
			localizedMessage := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(gameplayers.GamePlayerAlreadyRegisteredLexeme)(ctx),
			}
			st, _ = st.WithDetails(errorInfo, localizedMessage)
		} else if errors.Is(err, games.ErrGameNotFound) {
			st = status.New(codes.InvalidArgument, games.ErrGameNotFound.Error())
			errorInfo := &errdetails.ErrorInfo{
				Reason: games.ReasonGameNotFound,
				Metadata: map[string]string{
					"error": err.Error(),
				},
			}
			localizedMessage := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(games.GameNotFoundLexeme)(ctx),
			}
			st, _ = st.WithDetails(errorInfo, localizedMessage)
		} else if errors.Is(err, users.ErrUserNotFound) {
			st = status.New(codes.InvalidArgument, users.ErrUserNotFound.Error())
			errorInfo := &errdetails.ErrorInfo{
				Reason: users.ReasonUserNotFound,
				Metadata: map[string]string{
					"error": err.Error(),
				},
			}
			localizedMessage := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: i18n.GetTranslator(users.UserNotFoundLexeme)(ctx),
			}
			st, _ = st.WithDetails(errorInfo, localizedMessage)
		}

		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}

func validateRegisteredGamePlayer(gamePlayer model.GamePlayer) error {
	return validateCreatedGamePlayer(gamePlayer)
}
