package gameresultmanager

import (
	"context"

	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListGameResults ...
func (m *GameResultManager) ListGameResults(ctx context.Context, _ *emptypb.Empty) (*gameresultmanagerpb.ListGameResultsResponse, error) {
	gameResults, err := m.gameResultsFacade.ListGameResults(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	respGameResults := make([]*gameresultmanagerpb.GameResult, 0, len(gameResults))
	for _, gameResult := range gameResults {
		respGameResults = append(respGameResults, &gameresultmanagerpb.GameResult{
			Id:          gameResult.ID,
			GameId:      gameResult.FkGameID,
			ResultPlace: gameResult.ResultPlace,
			RoundPoints: gameResult.RoundPoints.Value,
		})
	}

	return &gameresultmanagerpb.ListGameResultsResponse{
		GameResults: respGameResults,
	}, nil
}
