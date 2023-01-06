package quiz_please

import "net/http"

const (
	lotteryLink = "https://quizplease.ru/ajax/send-lottery-player-form"
)

// QuizPleaseCroupier ...
type QuizPleaseCroupier struct {
	client      http.Client
	lotteryLink string
}

// New ...
func New() *QuizPleaseCroupier {
	return &QuizPleaseCroupier{
		client:      *http.DefaultClient,
		lotteryLink: lotteryLink,
	}
}
