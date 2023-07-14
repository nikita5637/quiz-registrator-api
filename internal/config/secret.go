package config

import "os"

const (
	// DatabasePassword ...
	DatabasePassword = "DATABASE_PASSWORD" // nolint:gosec
	// RabbitMQPassword ...
	RabbitMQPassword = "RABBITMQ_PASSWORD" // nolint:gosec
)

// GetSecretValue ...
func GetSecretValue(key string) string {
	return os.Getenv(key)
}
