package config

import (
	"github.com/spf13/viper"
)

func GetStringOrDefault(key, defaultVal string) string {
	str := Viper.GetString(key)
	if str == "" {
		return defaultVal
	}
	return str
}

func Default(key string, value interface{}) {
	Viper.SetDefault(key, value)
}

func Sub(key string) *viper.Viper { return Viper.Sub(key) }

func GetString(key string) string { return Viper.GetString(key) }

func GetStringMap(key string) map[string]interface{} {
	return Viper.GetStringMap(key)
}
func GetInt64(key string) int64 { return Viper.GetInt64(key) }
func GetUint(key string) uint   { return Viper.GetUint(key) }
func GetInt(key string) int     { return Viper.GetInt(key) }
func GetBool(key string) bool   { return Viper.GetBool(key) }
