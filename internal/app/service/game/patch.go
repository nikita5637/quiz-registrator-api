package game

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/leagues"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/places"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PatchGame ...
func (i *Implementation) PatchGame(ctx context.Context, req *gamepb.PatchGameRequest) (*gamepb.Game, error) {
	if req.GetGame() == nil {
		st := status.New(codes.InvalidArgument, "bad request")
		return nil, st.Err()
	}

	originalGame, err := i.gamesFacade.GetGameByID(ctx, req.GetGame().GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	patchedGame := originalGame
	for _, path := range req.GetUpdateMask().GetPaths() {
		switch path {
		case "external_id":
			if externalID := req.GetGame().GetExternalId(); externalID != nil {
				patchedGame.ExternalID = maybe.Just(externalID.GetValue())
			} else {
				patchedGame.ExternalID = maybe.Nothing[int32]()
			}
		case "league_id":
			patchedGame.LeagueID = req.GetGame().GetLeagueId()
		case "type":
			patchedGame.Type = model.GameType(req.GetGame().GetType())
		case "number":
			patchedGame.Number = req.GetGame().GetNumber()
		case "name":
			if name := req.GetGame().GetName(); name != nil {
				patchedGame.Name = maybe.Just(name.GetValue())
			} else {
				patchedGame.Name = maybe.Nothing[string]()
			}
		case "place_id":
			patchedGame.PlaceID = req.GetGame().GetPlaceId()
		case "date":
			patchedGame.Date = model.DateTime(req.GetGame().GetDate().AsTime())
		case "price":
			patchedGame.Price = req.GetGame().GetPrice()
		case "payment_type":
			if paymentType := req.GetGame().GetPaymentType(); paymentType != nil {
				patchedGame.PaymentType = maybe.Just(paymentType.GetValue())
			} else {
				patchedGame.PaymentType = maybe.Nothing[string]()
			}
		case "max_players":
			patchedGame.MaxPlayers = req.GetGame().GetMaxPlayers()
		case "payment":
			if payment := req.GetGame().Payment; payment != nil {
				patchedGame.Payment = maybe.Just(model.Payment(*payment))
			} else {
				patchedGame.Payment = maybe.Nothing[model.Payment]()
			}
		case "registered":
			patchedGame.Registered = req.GetGame().GetRegistered()
		case "is_in_master":
			patchedGame.IsInMaster = req.GetGame().GetIsInMaster()
		}
	}
	if err = validatePatchedGame(patchedGame); err != nil {
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

	game, err := i.gamesFacade.PatchGame(ctx, patchedGame)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, games.ErrGameAlreadyExists.Error(), games.ReasonGameAlreadyExists, map[string]string{
				"error": err.Error(),
			}, games.GameAlreadyExistsLexeme)
		} else if errors.Is(err, leagues.ErrLeagueNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, leagues.ErrLeagueNotFound.Error(), leagues.ReasonLeagueNotFound, map[string]string{
				"error": err.Error(),
			}, leagues.LeagueNotFoundLexeme)
		} else if errors.Is(err, places.ErrPlaceNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, places.ErrPlaceNotFound.Error(), places.ReasonPlaceNotFound, map[string]string{
				"error": err.Error(),
			}, places.PlaceNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameToProtoGame(game), nil
}

func validatePatchedGame(game model.Game) error {
	return validation.ValidateStruct(&game,
		validation.Field(&game.ExternalID, validation.By(validateExternalID)),
		validation.Field(&game.LeagueID, validation.Required, validation.Min(int32(1))),
		validation.Field(&game.Type, validation.Required, validation.By(model.ValidateGameType)),
		validation.Field(&game.Number, validation.When(game.Type != model.GameTypeClosed, validation.Required, validation.Length(1, 32)).Else(validation.Length(0, 32))),
		validation.Field(&game.Name, validation.By(validateName)),
		validation.Field(&game.PlaceID, validation.Required, validation.Min(int32(1))),
		validation.Field(&game.Date, validation.Required, validation.By(model.ValidateDateTime)),
		validation.Field(&game.PaymentType, validation.By(validatePaymentType)),
		validation.Field(&game.MaxPlayers, validation.Required, validation.Min(uint32(1))),
		validation.Field(&game.Payment, validation.By(validatePayment)),
	)

}
