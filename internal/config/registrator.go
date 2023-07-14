package config

import "fmt"

// RegistratorConfig ...
type RegistratorConfig struct {
	ActiveGameLag        uint16 `toml:"active_game_lag"`
	BindAddress          string `toml:"bind_address"`
	BindPort             uint16 `toml:"bind_port"`
	LotteryStartsBefore  uint16 `toml:"lottery_starts_before"`
	RabbitMQICSQueueName string `toml:"rabbitmq_ics_queue_name"`
}

// GetBindAddress ...
func GetBindAddress() string {
	return fmt.Sprintf("%s:%d", globalConfig.BindAddress, globalConfig.BindPort)
}
