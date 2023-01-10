package mysql

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// GameResultStorageAdapter ...
type GameResultStorageAdapter struct {
	gameResultStorage *GameResultStorage
}

// NewGameResultStorageAdapter ...
func NewGameResultStorageAdapter(txManager *tx.Manager) *GameResultStorageAdapter {
	return &GameResultStorageAdapter{
		gameResultStorage: NewGameResultStorage(txManager),
	}
}

// GetGameResultByFkGameID ...
func (a *GameResultStorageAdapter) GetGameResultByFkGameID(ctx context.Context, gameID int32) (model.GameResult, error) {
	dbGameResult, err := a.gameResultStorage.GetGameResultByFkGameID(ctx, int(gameID))
	if err != nil {
		return model.GameResult{}, err
	}

	return convertDBGameResultToModelGameResult(*dbGameResult), nil
}

func convertDBGameResultToModelGameResult(game GameResult) model.GameResult {
	return model.GameResult{
		ID:       int32(game.ID),
		FkGameID: int32(game.FkGameID),
		Place:    uint32(game.Place),
	}
}
