package mysql

import (
	"context"

	"github.com/mono83/maybe"
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

// CreateGameResult ...
func (a *GameResultStorageAdapter) CreateGameResult(ctx context.Context, dbGameResult GameResult) (int, error) {
	id, err := a.gameResultStorage.Insert(ctx, dbGameResult)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// GetGameResultByID ...
func (a *GameResultStorageAdapter) GetGameResultByID(ctx context.Context, id int) (*GameResult, error) {
	return a.gameResultStorage.GetGameResultByID(ctx, id)
}

// GetGameResults ...
func (a *GameResultStorageAdapter) GetGameResults(ctx context.Context) ([]GameResult, error) {
	return a.gameResultStorage.Find(ctx, nil, "id")
}

// GetGameResultByFkGameID ...
func (a *GameResultStorageAdapter) GetGameResultByFkGameID(ctx context.Context, gameID int) (model.GameResult, error) {
	dbGameResult, err := a.gameResultStorage.GetGameResultByFkGameID(ctx, gameID)
	if err != nil {
		return model.GameResult{}, err
	}

	return convertDBGameResultToModelGameResult(*dbGameResult), nil
}

// PatchGameResult ...
func (a *GameResultStorageAdapter) PatchGameResult(ctx context.Context, dbGameResult GameResult) error {
	return a.gameResultStorage.Update(ctx, dbGameResult)
}

func convertDBGameResultToModelGameResult(game GameResult) model.GameResult {
	ret := model.GameResult{
		ID:          int32(game.ID),
		FkGameID:    int32(game.FkGameID),
		ResultPlace: uint32(game.Place),
		RoundPoints: maybe.Nothing[string](),
	}

	if game.Points.Valid {
		ret.RoundPoints = maybe.Just(game.Points.String)
	}

	return ret
}
