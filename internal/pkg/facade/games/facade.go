//go:generate mockery --case underscore --name QuizLogger --with-expecter

package games

import (
	"context"

	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// QuizLogger ...
type QuizLogger interface {
	Write(ctx context.Context, params quizlogger.Params) error
	WriteBatch(ctx context.Context, params []quizlogger.Params) error
}

// Facade ...
type Facade struct {
	db          *tx.Manager
	gameStorage storage.GameStorage
	quizLogger  QuizLogger
}

// Config ...
type Config struct {
	GameStorage storage.GameStorage
	TxManager   *tx.Manager
	QuizLogger  QuizLogger
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:          cfg.TxManager,
		gameStorage: cfg.GameStorage,
		quizLogger:  cfg.QuizLogger,
	}
}
