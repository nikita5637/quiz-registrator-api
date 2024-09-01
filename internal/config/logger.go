package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	defaultLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func initLoggerConfigureParams() {
	_ = viper.BindEnv("log.elastic.addresses")
}

// GetLogLevel ...
func GetLogLevel() zap.AtomicLevel {
	level, err := zap.ParseAtomicLevel(viper.GetString("log.level"))
	if err != nil {
		level = defaultLogLevel
	}

	return level
}
