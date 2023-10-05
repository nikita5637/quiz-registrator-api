package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// amqp://guest:guest@localhost:5672/
	rabbitMQURLFormat = "amqp://%s:%s@%s:%d/"
)

func initRemindManagerConfigureParams() {
	_ = viper.BindEnv("remind_manager.rabbitmq.address")
	_ = viper.BindEnv("remind_manager.rabbitmq.credentials.password")
}

// GetRabbitMQURL ...
func GetRabbitMQURL() string {
	return fmt.Sprintf(rabbitMQURLFormat,
		viper.GetString("remind_manager.rabbitmq.credentials.username"),
		viper.GetString("remind_manager.rabbitmq.credentials.password"),
		viper.GetString("remind_manager.rabbitmq.address"),
		viper.GetUint32("remind_manager.rabbitmq.port"),
	)
}
