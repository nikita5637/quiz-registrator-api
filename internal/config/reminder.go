package config

import (
	"fmt"
)

const (
	// amqp://guest:guest@localhost:5672/
	rabbitMQURLFormat = "amqp://%s:%s@%s:%d/"
)

// ReminderConfig ...
type ReminderConfig struct {
	RabbitMQAddress                  string `toml:"rabbitmq_address"`
	RabbitMQGameReminderQueueName    string `toml:"rabbitmq_game_reminder_queue_name"`
	RabbitMQLotteryReminderQueueName string `toml:"rabbitmq_lottery_reminder_queue_name"`
	RabbitMQPort                     uint16 `toml:"rabbitmq_port"`
	RabbitMQUserName                 string `toml:"rabbitmq_username"`
}

// GetRabbitMQURL ...
func GetRabbitMQURL() string {
	rabbitMQPassword := GetSecretValue(RabbitMQPassword)
	return fmt.Sprintf(rabbitMQURLFormat,
		globalConfig.RabbitMQUserName,
		rabbitMQPassword,
		globalConfig.RabbitMQAddress,
		globalConfig.RabbitMQPort,
	)
}
