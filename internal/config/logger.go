package config

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	defaultLogLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func initLoggerConfigureParams() {
	_ = viper.BindEnv("log.elastic.address")
}

// GetElasticAddress ...
func GetElasticAddress() string {
	return fmt.Sprintf("http://%s:%d",
		viper.GetString("log.elastic.address"),
		viper.GetUint64("log.elastic.port"),
	)
}

// GetLogLevel ...
func GetLogLevel() zap.AtomicLevel {
	level, err := zap.ParseAtomicLevel(viper.GetString("log.level"))
	if err != nil {
		level = defaultLogLevel
	}

	return level
}
