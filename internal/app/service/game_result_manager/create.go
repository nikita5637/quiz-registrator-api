package gameresultmanager

import (
	"context"
	"errors"
	"fmt"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateGameResult ...
func (m *GameResultManager) CreateGameResult(ctx context.Context, req *gameresultmanagerpb.CreateGameResultRequest) (*gameresultmanagerpb.GameResult, error) {
	if err := validateCreateGameResultRequest(ctx, req); err != nil {
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

	gameResult, err := m.gameResultsFacade.CreateGameResult(ctx, model.GameResult{
		FkGameID:    req.GetGameResult().GetGameId(),
		ResultPlace: req.GetGameResult().GetResultPlace(),
		RoundPoints: maybe.Just(req.GetGameResult().GetRoundPoints()),
	})
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, model.ErrGameNotFound) {
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

func validateCreateGameResultRequest(ctx context.Context, req *gameresultmanagerpb.CreateGameResultRequest) error {
	return validateGameResult(ctx, req.GetGameResult())
}
