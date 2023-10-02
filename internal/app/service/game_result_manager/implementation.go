//go:generate mockery --case underscore --name GameResultsFacade --with-expecter

package gameresultmanager

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameresultmanagerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_result_manager"
)

// GameResultsFacade ...
type GameResultsFacade interface {
	CreateGameResult(ctx context.Context, gameResult model.GameResult) (model.GameResult, error)
	ListGameResults(ctx context.Context) ([]model.GameResult, error)
	PatchGameResult(ctx context.Context, gameResult model.GameResult, paths []string) (model.GameResult, error)
	SearchGameResultByGameID(ctx context.Context, gameID int32) (model.GameResult, error)
}

// GameResultManager ...
type GameResultManager struct {
	gameResultsFacade GameResultsFacade

	gameresultmanagerpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	GameResultsFacade GameResultsFacade
}

// New ...
func New(cfg Config) *GameResultManager {
	return &GameResultManager{
		gameResultsFacade: cfg.GameResultsFacade,
	}
}
