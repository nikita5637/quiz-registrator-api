//go:generate mockery --case underscore --name Croupier --with-expecter
//go:generate mockery --case underscore --name GamePlayersFacade --with-expecter
//go:generate mockery --case underscore --name GamesFacade --with-expecter

package croupier

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	croupierpb "github.com/nikita5637/quiz-registrator-api/pkg/pb/croupier"
)

// Croupier ...
type Croupier interface {
	GetIsLotteryActive(ctx context.Context, game model.Game) bool
	RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error)
}

// GamesFacade ...
type GamesFacade interface {
	GetGameByID(ctx context.Context, id int32) (model.Game, error)
}

// GamePlayersFacade ...
type GamePlayersFacade interface {
	PlayerRegisteredOnGame(ctx context.Context, gameID, userID int32) (bool, error)
}

// Implemintation ...
type Implemintation struct {
	croupier          Croupier
	gamePlayersFacade GamePlayersFacade
	gamesFacade       GamesFacade

	croupierpb.UnimplementedServiceServer
}

// Config ...
type Config struct {
	Croupier          Croupier
	GamePlayersFacade GamePlayersFacade
	GamesFacade       GamesFacade
}

// New ...
func New(cfg Config) *Implemintation {
	return &Implemintation{
		croupier:          cfg.Croupier,
		gamePlayersFacade: cfg.GamePlayersFacade,
		gamesFacade:       cfg.GamesFacade,
	}
}
