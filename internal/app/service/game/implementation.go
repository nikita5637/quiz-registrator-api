//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name RabbitMQProducer --with-expecter

package game

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
)

// GamesFacade ...
type GamesFacade interface {
	CreateGame(ctx context.Context, game model.Game) (model.Game, error)
	DeleteGame(ctx context.Context, id int32) error
	GetGameByID(ctx context.Context, id int32) (model.Game, error)
	GetGamesByIDs(ctx context.Context, ids []int32) ([]model.Game, error)
	ListGames(ctx context.Context) ([]model.Game, error)
	PatchGame(ctx context.Context, game model.Game) (model.Game, error)
	SearchGamesByLeagueID(ctx context.Context, leagueID int32, offset, limit uint64) ([]model.Game, uint64, error)
	SearchPassedAndRegisteredGames(ctx context.Context, offset, limit uint64) ([]model.Game, uint64, error)
}

// RabbitMQProducer ...
type RabbitMQProducer interface {
	Send(ctx context.Context, message interface{}) error
}

// Implementation ...
type Implementation struct {
	gamesFacade      GamesFacade
	rabbitMQProducer RabbitMQProducer

	gamepb.UnimplementedServiceServer
	gamepb.UnimplementedRegistratorServiceServer
}

// Config ...
type Config struct {
	GamesFacade      GamesFacade
	RabbitMQProducer RabbitMQProducer
}

// New ...
func New(cfg Config) *Implementation {
	return &Implementation{
		gamesFacade:      cfg.GamesFacade,
		rabbitMQProducer: cfg.RabbitMQProducer,
	}
}
