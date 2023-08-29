//go:generate mockery --case underscore --name GamesFacade --with-expecter

package registrator

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	registratorpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// GamesFacade ...
type GamesFacade interface {
	AddGame(ctx context.Context, game model.Game) (int32, error)
	AddGames(ctx context.Context, games []model.Game) error
	DeleteGame(ctx context.Context, gameID int32) error
	// GetGameByID guaranteed returns active game by game ID
	GetGameByID(ctx context.Context, id int32) (model.Game, error)
	GetGames(ctx context.Context) ([]model.Game, error)
	GetGamesByUserID(ctx context.Context, userID int32) ([]model.Game, error)
	GetRegisteredGames(ctx context.Context) ([]model.Game, error)
	RegisterGame(ctx context.Context, gameID int32) (model.RegisterGameStatus, error)
	UnregisterGame(ctx context.Context, gameID int32) (model.UnregisterGameStatus, error)
	UpdatePayment(ctx context.Context, gameID int32, payment model.PaymentType) error
}

// Implementation ...
type Implementation struct {
	gamesFacade GamesFacade

	registratorpb.UnimplementedRegistratorServiceServer
}

// Config ...
type Config struct {
	GamesFacade GamesFacade
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		gamesFacade: cfg.GamesFacade,
	}
}
