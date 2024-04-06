package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initAPIServerConfigureParams() {
	_ = viper.BindEnv("apiserver.grpc.bind.address")
}

// GetGRPCBindAddress ...
func GetGRPCBindAddress() string {
	return fmt.Sprintf("%s:%d",
		viper.GetString("apiserver.grpc.bind.address"),
		viper.GetUint32("apiserver.grpc.bind.port"),
	)
}
