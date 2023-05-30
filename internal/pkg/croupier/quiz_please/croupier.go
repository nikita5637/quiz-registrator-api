package quiz_please

import "net/http"

const (
	// LotteryLink ...
	LotteryLink = "https://quizplease.ru/ajax/send-lottery-player-form"
)

// Croupier ...
type Croupier struct {
	client      http.Client
	lotteryLink string
}

// Config ...
type Config struct {
	LotteryLink string
}

// New ...
func New(cfg Config) *Croupier {
	return &Croupier{
		client:      *http.DefaultClient,
		lotteryLink: cfg.LotteryLink,
	}
}
