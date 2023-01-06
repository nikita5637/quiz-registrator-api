package telegram

import "context"

type telegramClientIDKeyType struct{}

var (
	telegramClientIDKey = telegramClientIDKeyType{}
)

// NewContextWithClientID ...
func NewContextWithClientID(ctx context.Context, clientID int64) context.Context {
	return context.WithValue(ctx, telegramClientIDKey, clientID)
}

// ClientIDFromContext ...
func ClientIDFromContext(ctx context.Context) int64 {
	val, ok := ctx.Value(telegramClientIDKey).(int64)
	if !ok {
		return 0
	}

	return val
}
