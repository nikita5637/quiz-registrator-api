package squiz

import "net/http"

const (
	// LotteryInfoPageLink ...
	LotteryInfoPageLink = "https://spb.squiz.ru/game"
	// LotteryRegistrationLink ...
	LotteryRegistrationLink = "https://forms.tildacdn.com/procces"
)

// Croupier ...
type Croupier struct {
	client                  http.Client
	lotteryInfoPageLink     string
	lotteryRegistrationLink string
}

// Config ...
type Config struct {
	LotteryInfoPageLink     string
	LotteryRegistrationLink string
}

// New ...
func New(cfg Config) *Croupier {
	return &Croupier{
		client:                  *http.DefaultClient,
		lotteryInfoPageLink:     cfg.LotteryInfoPageLink,
		lotteryRegistrationLink: cfg.LotteryRegistrationLink,
	}
}
