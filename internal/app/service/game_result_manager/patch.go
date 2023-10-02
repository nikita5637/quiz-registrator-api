package gameresultmanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameresults"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/games"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PatchGameResult ...
func (m *GameResultManager) PatchGameResult(ctx context.Context, req *gameresultmanagerpb.PatchGameResultRequest) (*gameresultmanagerpb.GameResult, error) {
	if err := validatePatchGameResultRequest(ctx, req); err != nil {
		st := status.New(codes.InvalidArgument, err.Error())
		if errors.Is(err, errInvalidJSONRoundPointsValue) {
			reason := fmt.Sprintf("invalid game result round points JSON value: \"%s\"", req.GetGameResult().GetRoundPoints())
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), reason, nil, invalidGameResultRoundPointsJSONValueLexeme)
		} else if errors.Is(err, errInvalidResultPlace) {
			reason := fmt.Sprintf("invalid game result result place: \"%d\"", req.GetGameResult().GetResultPlace())
			st = model.GetStatus(ctx, codes.InvalidArgument, err.Error(), reason, nil, invalidGameResultResultPlaceValueLexeme)
		}

		return nil, st.Err()
	}

	gameResult, err := m.gameResultsFacade.PatchGameResult(ctx, model.GameResult{
		ID:          req.GetGameResult().GetId(),
		FkGameID:    req.GetGameResult().GetGameId(),
		ResultPlace: req.GetGameResult().GetResultPlace(),
		RoundPoints: maybe.Just(req.GetGameResult().GetRoundPoints()),
	}, req.GetUpdateMask().GetPaths())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameresults.ErrGameResultNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, gameresults.ErrGameResultNotFound.Error(), gameresults.ReasonGameResultNotFound, map[string]string{
				"error": err.Error(),
			}, gameresults.GameResultNotFoundLexeme)
		} else if errors.Is(err, games.ErrGameNotFound) {
			st = model.GetStatus(ctx, codes.FailedPrecondition, games.ErrGameNotFound.Error(), games.ReasonGameNotFound, map[string]string{
				"error": err.Error(),
			}, games.GameNotFoundLexeme)
		} else if errors.Is(err, gameresults.ErrGameResultAlreadyExists) {
			st = model.GetStatus(ctx, codes.AlreadyExists, gameresults.ErrGameResultAlreadyExists.Error(), gameresults.ReasonGameResultAlreadyExists, map[string]string{
				"error": err.Error(),
			}, gameresults.GameResultAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameResultToProtoGameResult(gameResult), nil
}

func validatePatchGameResultRequest(ctx context.Context, req *gameresultmanagerpb.PatchGameResultRequest) error {
	return validateGameResult(ctx, req.GetGameResult())
}
