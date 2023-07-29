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
)

// CreateGamePlayer ...
func (i *Implementation) CreateGamePlayer(ctx context.Context, req *gameplayerpb.CreateGamePlayerRequest) (*gameplayerpb.GamePlayer, error) {
	if req.GetGamePlayer() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	createdGamePlayer := convertProtoGamePlayerToModelGamePlayer(req.GetGamePlayer())

	logger.Debugf(ctx, "trying to create new game player: %#v", createdGamePlayer)

	if err := validateCreatedGamePlayer(createdGamePlayer); err != nil {
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

	gamePlayer, err := i.gamePlayersFacade.CreateGamePlayer(ctx, createdGamePlayer)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameplayers.ErrGamePlayerAlreadyRegistered) {
			st = model.GetStatus(ctx, codes.AlreadyExists, gameplayers.ErrGamePlayerAlreadyRegistered, gameplayers.ReasonGamePlayerAlreadyRegistered, gameplayers.GamePlayerAlreadyRegisteredLexeme)
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

	return convertModelGamePlayerToProtoGamePlayer(gamePlayer), nil
}

func validateCreatedGamePlayer(gamePlayer model.GamePlayer) error {
	return validation.ValidateStruct(&gamePlayer,
		validation.Field(&gamePlayer.GameID, validation.Required),
		validation.Field(&gamePlayer.UserID, validation.By(validateUserID)),
		validation.Field(&gamePlayer.RegisteredBy, validation.Required),
		validation.Field(&gamePlayer.Degree, validation.Required, validation.By(model.ValidateDegree)),
	)
}
