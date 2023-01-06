package config

import "os"

const (
	// DatabasePassword ...
	DatabasePassword = "DATABASE_PASSWORD" // nolint:gosec
	// TelegramToken ...
	TelegramToken = "QUIZ_TELEGRAM_BOT_TOKEN" // nolint:gosec
)

// GetSecretValue ...
func GetSecretValue(key string) string {
	return os.Getenv(key)
}
