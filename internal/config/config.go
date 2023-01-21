package config

import (
	"reflect"

	"github.com/BurntSushi/toml"
)

// GlobalConfig ...
type GlobalConfig struct {
	DatabaseConfig
	LoggerConfig
	RegistratorConfig
	ReminderConfig
}

var globalConfig GlobalConfig

// Value ...
type Value reflect.Value

// GetValue ...
func GetValue(key string) Value {
	val := reflect.Indirect(reflect.ValueOf(globalConfig)).FieldByName(key)
	return Value(val)
}

// Bool ...
func (c Value) Bool() bool {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return false
	}

	return reflect.Value(c).Bool()
}

// String ...
func (c Value) String() string {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return ""
	}

	return reflect.Value(c).String()
}

// Uint16 ...
func (c Value) Uint16() uint16 {
	if reflect.Value(c).Kind() == reflect.Invalid {
		return 0
	}

	return uint16(reflect.Value(c).Uint())
}

// ParseConfigFile ...
func ParseConfigFile(path string) error {
	_, err := toml.DecodeFile(path, &globalConfig)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGlobalConfig ...
func UpdateGlobalConfig(newGlobalConfig GlobalConfig) {
	globalConfig = newGlobalConfig
}
