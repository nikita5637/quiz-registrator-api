package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initAPIServerConfigureParams() {
	_ = viper.BindEnv("apiserver.bind.address")
}

// GetBindAddress ...
func GetBindAddress() string {
	return fmt.Sprintf("%s:%d",
		viper.GetString("apiserver.bind.address"),
		viper.GetUint32("apiserver.bind.port"),
	)
}
