package gameresultmanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	gameResultNotFoundLexeme = i18n.Lexeme{
		Key:      "game_result_not_found_lexeme",
		FallBack: "Game result not found",
	}
)

// PatchGameResult ...
func (m *GameResultManager) PatchGameResult(ctx context.Context, req *gameresultmanagerpb.PatchGameResultRequest) (*gameresultmanagerpb.GameResult, error) {
	if err := validatePatchGameResultRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONRoundPointsValue) {
			reason := fmt.Sprintf("invalid game result round points JSON value: \"%s\"", req.GetGameResult().GetRoundPoints())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultRoundPointsJSONValueLexeme)
		} else if errors.Is(err, errInvalidResultPlace) {
			reason := fmt.Sprintf("invalid game result result place: \"%d\"", req.GetGameResult().GetResultPlace())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, invalidGameResultResultPlaceValueLexeme)
		}

		return nil, st.Err()
	}

	gameResult, err := m.gameResultsFacade.PatchGameResult(ctx, model.GameResult{
		ID:          req.GetGameResult().GetId(),
		FkGameID:    req.GetGameResult().GetGameId(),
		ResultPlace: req.GetGameResult().GetResultPlace(),
		RoundPoints: model.NewMaybeString(req.GetGameResult().GetRoundPoints()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrGameResultNotFound) {
			reason := fmt.Sprintf("game result with ID %d not found", req.GetGameResult().GetId())
			st = model.GetStatus(ctx, codes.NotFound, err, reason, gameResultNotFoundLexeme)
		} else if errors.Is(err, model.ErrGameNotFound) {
			reason := fmt.Sprintf("game with id %d not found", req.GetGameResult().GetGameId())
			st = model.GetStatus(ctx, codes.InvalidArgument, err, reason, i18n.GameNotFoundLexeme)
		} else if errors.Is(err, model.ErrGameResultAlreadyExists) {
			reason := fmt.Sprintf("game result for game id %d already exists", req.GetGameResult().GetGameId())
			st = model.GetStatus(ctx, codes.AlreadyExists, err, reason, gameResultAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameResultToProtoGameResult(gameResult), nil
}

func validatePatchGameResultRequest(ctx context.Context, req *gameresultmanagerpb.PatchGameResultRequest) error {
	return validateGameResult(ctx, req.GetGameResult())
}
