//go:generate mockery --case underscore --name QuizLogger --with-expecter

package certificates

import (
	"context"

	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
)

// QuizLogger ...
type QuizLogger interface {
	Write(ctx context.Context, params quizlogger.Params) error
}

// Facade ...
type Facade struct {
	db                 *tx.Manager
	certificateStorage storage.CertificateStorage
	quizLogger         QuizLogger
}

// Config ...
type Config struct {
	CertificateStorage storage.CertificateStorage
	TxManager          *tx.Manager
	QuizLogger         QuizLogger
}

// New ...
func New(cfg Config) *Facade {
	return &Facade{
		db:                 cfg.TxManager,
		certificateStorage: cfg.CertificateStorage,
		quizLogger:         cfg.QuizLogger,
	}
}
