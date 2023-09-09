package croupier

import (
	"context"
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	lotteryNotAvailableLexeme = i18n.Lexeme{
		Key:      "lottery_not_available",
		FallBack: "Lottery not available",
	}
	lotteryNotImplementedLexeme = i18n.Lexeme{
		Key:      "lottery_not_implemented",
		FallBack: "Lottery not implemented for this league",
	}
	lotteryPermissionDenied = i18n.Lexeme{
		Key:      "lottery_permission_denied",
		FallBack: "Lottery permission denied",
	}
)

// RegisterForLottery ...
func (i *Implemintation) RegisterForLottery(ctx context.Context, req *croupierpb.RegisterForLotteryRequest) (*croupierpb.RegisterForLotteryResponse, error) {
	if err := validateRegisterForLotteryRequest(req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		return nil, st.Err()
	}

	user := usersutils.UserFromContext(ctx)
	if user.Email.Value() == "" || user.Name == "" || user.Phone.Value() == "" {
		reason := fmt.Sprintf("permission denied for lottery registration for user %d", user.ID)
		st := model.GetStatus(ctx, codes.PermissionDenied, "permission denied for lottery registration", reason, nil, lotteryPermissionDenied)
		return nil, st.Err()
	}

	userRegistered, err := i.gamePlayersFacade.PlayerRegisteredOnGame(ctx, req.GetGameId(), user.ID)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	if !userRegistered {
		reason := fmt.Sprintf("permission denied for lottery registration for user %d", user.ID)
		st := model.GetStatus(ctx, codes.PermissionDenied, "permission denied for lottery registration", reason, nil, lotteryPermissionDenied)
		return nil, st.Err()
	}

	game, err := i.gamesFacade.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, games.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		} else if errors.Is(err, games.ErrGameHasPassed) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}
		return nil, st.Err()
	}

	var number int32
	number, err = i.croupier.RegisterForLottery(ctx, game, *user)
	if err != nil {
		st := status.New(codes.InvalidArgument, err.Error())

		if errors.Is(err, model.ErrLotteryNotAvailable) {
			reason := fmt.Sprintf("lottery for game id %d not available", game.ID)
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), reason, nil, lotteryNotAvailableLexeme)
		} else if errors.Is(err, model.ErrLotteryNotImplemented) {
			reason := fmt.Sprintf("lottery for league %d not implemented", game.LeagueID)
			st = model.GetStatus(ctx, codes.Unimplemented, err.Error(), reason, nil, lotteryNotImplementedLexeme)
		} else {
			if unwrappedError := errors.Unwrap(err); unwrappedError != nil {
				err = unwrappedError
			}

			ei := &errdetails.ErrorInfo{
				Reason: fmt.Sprintf("lottery registration for game ID %d for user ID %d failed", game.ID, user.ID),
			}
			lm := &errdetails.LocalizedMessage{
				Locale:  i18n.GetLangFromContext(ctx),
				Message: err.Error(),
			}
			st, err = st.WithDetails(ei, lm)
			if err != nil {
				panic(err)
			}
		}

		return nil, st.Err()
	}

	return &croupierpb.RegisterForLotteryResponse{
		Number: number,
	}, nil
}

func validateRegisterForLotteryRequest(req *croupierpb.RegisterForLotteryRequest) error {
	err := validation.Validate(req.GetGameId(), validation.Required, validation.Min(int32(1)))
	if err != nil {
		return err
	}

	return nil
}
