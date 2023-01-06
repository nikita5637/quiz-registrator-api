package mysql

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// GameResultStorageAdapter ...
type GameResultStorageAdapter struct {
	gameResultStorage *GameResultStorage
}

// NewGameResultStorageAdapter ...
func NewGameResultStorageAdapter(db *sql.DB) *GameResultStorageAdapter {
	return &GameResultStorageAdapter{
		gameResultStorage: NewGameResultStorage(db),
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
