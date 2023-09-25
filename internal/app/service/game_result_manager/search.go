package gameresultmanager

import (
	"context"
	"errors"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/facade/gameresults"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SearchGameResultByGameID ...
func (m *GameResultManager) SearchGameResultByGameID(ctx context.Context, req *gameresultmanagerpb.SearchGameResultByGameIDRequest) (*gameresultmanagerpb.GameResult, error) {
	gameResult, err := m.gameResultsFacade.SearchGameResultByGameID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		if errors.Is(err, gameresults.ErrGameResultNotFound) {
			st = model.GetStatus(ctx, codes.NotFound, gameresults.ErrGameResultNotFound.Error(), gameresults.ReasonGameResultNotFound, map[string]string{
				"error": err.Error(),
			}, gameresults.GameResultNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return convertModelGameResultToProtoGameResult(gameResult), nil
}
