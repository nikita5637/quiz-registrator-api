package config

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	initAPIServerConfigureParams()
	initDatabaseConfigureParams()
	initLoggerConfigureParams()
	initRemindManagerConfigureParams()
}

// ReadConfig ...
func ReadConfig() error {
	if viper.IsSet("config") {
		viper.SetConfigFile(viper.GetString("config"))
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		viper.WatchConfig()
		return nil
	}

	return errors.New("config file is not specified")
}
