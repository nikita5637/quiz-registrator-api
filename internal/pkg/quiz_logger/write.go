package quizlogger

import (
	"context"
	"fmt"

	"github.com/mono83/maybe"
)

// Params ...
type Params struct {
	UserID     maybe.Maybe[int32]
	ActionID   int
	MessageID  int
	ObjectType maybe.Maybe[string]
	ObjectID   maybe.Maybe[int32]
	Metadata   interface{}
}

// Write ...
func (l *Logger) Write(ctx context.Context, params Params) error {
	return l.logStorage.Create(ctx, convertParamsToDBLog(params))
}

// WriteBatch ...
func (l *Logger) WriteBatch(ctx context.Context, paramsBatch []Params) error {
	for _, params := range paramsBatch {
		if err := l.logStorage.Create(ctx, convertParamsToDBLog(params)); err != nil {
			return fmt.Errorf("creating log error: %w", err)
		}
	}

	return nil
}
