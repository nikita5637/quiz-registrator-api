//go:generate mockery --case underscore --name GamesFacade --with-expecter
//go:generate mockery --case underscore --name LotteryRegistrator --with-expecter

package croupier

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
)

// GamesFacade ...
type GamesFacade interface {
	ListGames(ctx context.Context) ([]model.Game, error)
}

// LotteryRegistrator ...
type LotteryRegistrator interface {
	RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error)
}

// Croupier ...
type Croupier struct {
	gamesFacade        GamesFacade
	leaguesWithLottery []int32
	quizPleaseCroupier LotteryRegistrator
	squizCroupier      LotteryRegistrator
}

// Config ...
type Config struct {
	GamesFacade        GamesFacade
	QuizPleaseCroupier LotteryRegistrator
	SquizCroupier      LotteryRegistrator
}

// New ...
func New(cfg Config) *Croupier {
	return &Croupier{
		gamesFacade: cfg.GamesFacade,
		leaguesWithLottery: []int32{
			int32(leaguepb.LeagueID_QUIZ_PLEASE),
			int32(leaguepb.LeagueID_SQUIZ),
		},
		quizPleaseCroupier: cfg.QuizPleaseCroupier,
		squizCroupier:      cfg.SquizCroupier,
	}
}
