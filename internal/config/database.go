package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	// username:password@protocol(ip:port)/dbname
	mysqlDSNFormat = "%s:%s@tcp(%s:%d)/%s?parseTime=true"
)

func initDatabaseConfigureParams() {
	_ = viper.BindEnv("database.address")
	_ = viper.BindEnv("database.credentials.password")
}

// GetMySQLDatabaseDSN ...
func GetMySQLDatabaseDSN() string {
	return fmt.Sprintf(mysqlDSNFormat,
		viper.GetString("database.credentials.username"),
		viper.GetString("database.credentials.password"),
		viper.GetString("database.address"),
		viper.GetUint32("database.port"),
		viper.GetString("database.dbname"),
	)
}
