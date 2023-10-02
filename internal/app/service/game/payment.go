package game

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UpdatePayment ...
func (i *Implementation) UpdatePayment(ctx context.Context, req *gamepb.UpdatePaymentRequest) (*emptypb.Empty, error) {
	game, err := i.gamesFacade.GetGameByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		}

		return nil, st.Err()
	}

	if game.HasPassed {
		st := model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameHasPassed.Error(), games.ReasonGameHasPassed, nil, games.GameHasPassedLexeme)
		return nil, st.Err()
	}

	game.Payment = maybe.Nothing[model.Payment]()
	if payment := req.GetPayment(); payment > 0 {
		game.Payment = maybe.Just(model.Payment(req.GetPayment()))
	}

	if err = validation.ValidateStruct(&game,
		validation.Field(&game.Payment, validation.By(validatePayment)),
	); err != nil {
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

	if _, err = i.gamesFacade.PatchGame(ctx, game); err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	return &emptypb.Empty{}, nil
}
