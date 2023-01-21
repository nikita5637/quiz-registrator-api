//go:generate mockery --case underscore --name Reminder --with-expecter

package remindmanager

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
)

// Reminder ...
type Reminder interface {
	Run(ctx context.Context) error
}

// Manager ...
type Manager struct {
	reminders []Reminder
}

// Config ...
type Config struct {
	Reminders []Reminder
}

// New ...
func New(cfg Config) *Manager {
	return &Manager{
		reminders: cfg.Reminders,
	}
}

// Start ...
func (r *Manager) Start(ctx context.Context) error {
	go func(ctx context.Context) {
		for range time.Tick(1 * time.Minute) {
			for _, reminder := range r.reminders {
				go func(ctx context.Context, reminder Reminder) {
					_ = reminder.Run(ctx)
				}(ctx, reminder)
			}
		}
	}(ctx)

	<-ctx.Done()

	logger.Info(ctx, "remind manager gracefully stopped")
	return nil
}
