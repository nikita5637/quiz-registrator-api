//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name GamePlayersFacade --with-expecter

package gameplayer

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gameplayerpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game_player"
)

// GamesFacade ...
type GamesFacade interface {
	GetGameByID(ctx context.Context, gameID int32) (model.Game, error)
}

// GamePlayersFacade ...
type GamePlayersFacade interface {
	CreateGamePlayer(ctx context.Context, gamePlayer model.GamePlayer) (model.GamePlayer, error)
	DeleteGamePlayer(ctx context.Context, id int32) error
	GetGamePlayer(ctx context.Context, id int32) (model.GamePlayer, error)
	GetGamePlayersByGameID(ctx context.Context, gameID int32) ([]model.GamePlayer, error)
	PatchGamePlayer(ctx context.Context, gamePlayer model.GamePlayer) (model.GamePlayer, error)
}

// Implementation ...
type Implementation struct {
	gamesFacade       GamesFacade
	gamePlayersFacade GamePlayersFacade

	gameplayerpb.UnimplementedServiceServer
	gameplayerpb.UnimplementedRegistratorServiceServer
}

// Config ...
type Config struct {
	GamesFacade       GamesFacade
	GamePlayersFacade GamePlayersFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		gamesFacade:       cfg.GamesFacade,
		gamePlayersFacade: cfg.GamePlayersFacade,
	}
}
