package quizlogger

import (
	"context"

	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
)

// LogStorage ...
type LogStorage interface {
	Create(ctx context.Context, log database.Log) error
}

// Logger ...
type Logger struct {
	logStorage LogStorage
}

// Config ...
type Config struct {
	LogStorage LogStorage
}

// New ...
func New(cfg Config) *Logger {
	return &Logger{
		logStorage: cfg.LogStorage,
	}
}
