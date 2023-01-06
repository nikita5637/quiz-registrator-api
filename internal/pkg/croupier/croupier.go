//go:generate mockery --case underscore --name LotteryRegistrator --with-expecter

package croupier

import (
	"context"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
)

// LotteryRegistrator ...
type LotteryRegistrator interface {
	RegisterForLottery(ctx context.Context, game model.Game, user model.User) (int32, error)
}

// Croupier ...
type Croupier struct {
	leaguesWithLottery []int32
	quizPleaseCroupier LotteryRegistrator
}

// Config ...
type Config struct {
	QuizPleaseCroupier LotteryRegistrator
}

// New ...
func New(cfg Config) *Croupier {
	return &Croupier{
		leaguesWithLottery: []int32{
			model.LeagueQuizPlease,
		},
		quizPleaseCroupier: cfg.QuizPleaseCroupier,
	}
}
