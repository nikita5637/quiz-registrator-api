package registrator

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

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

// GetLotteryStatus ...
func (r *Registrator) GetLotteryStatus(ctx context.Context, req *registrator.GetLotteryStatusRequest) (*registrator.GetLotteryStatusResponse, error) {
	game, err := r.gamesFacade.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	return &registrator.GetLotteryStatusResponse{
		Active: r.croupier.GetIsLotteryActive(ctx, game),
	}, nil
}

// RegisterForLottery ...
func (r *Registrator) RegisterForLottery(ctx context.Context, req *registrator.RegisterForLotteryRequest) (*registrator.RegisterForLotteryResponse, error) {
	user := users_utils.UserFromContext(ctx)

	game, err := r.gamesFacade.GetGameByID(ctx, req.GetGameId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrGameNotFound) {
			st = getGameNotFoundStatus(ctx, err, req.GetGameId())
		}

		return nil, st.Err()
	}

	var number int32
	number, err = r.croupier.RegisterForLottery(ctx, game, user)
	if err != nil {
		st := status.New(codes.InvalidArgument, err.Error())

		if errors.Is(err, model.ErrLotteryNotAvailable) {
			reason := fmt.Sprintf("lottery for game id %d not available", game.ID)
			st = getStatus(ctx, codes.InvalidArgument, err, reason, lotteryNotAvailableLexeme)
		} else if errors.Is(err, model.ErrLotteryNotImplemented) {
			reason := fmt.Sprintf("lottery for league %d not implemented", game.LeagueID)
			st = getStatus(ctx, codes.Unimplemented, err, reason, lotteryNotImplementedLexeme)
		} else if errors.Is(err, model.ErrLotteryPermissionDenied) {
			reason := fmt.Sprintf("permission denied for lottery registration for user %d", user.ID)
			st = getStatus(ctx, codes.InvalidArgument, err, reason, lotteryPermissionDenied)
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

	return &registrator.RegisterForLotteryResponse{
		Number: number,
	}, nil
}
