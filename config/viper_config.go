package config

import (
	"context"
	"errors"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/spf13/viper"
	"log"
)

func init() {
	event.RegisterHook(event.Start, InitConfig)
}
func InitConfig() {
	ReadLocalConfig()
	if viper.GetStringMap("cc.cloud.nacos.config") != nil {
		NacosInit()
	}
	event.TriggerEvent(event.ConfigComplete)
}
func ReadLocalConfig() {
	// 读取本地配置
	viper.SetConfigName("bootstrap.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")
	viper.AddConfigPath("/etc/resource")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info(context.Background(), "Config file not found; ignore error if desired")

		} else {
			logger.Info(context.Background(), "Config file was found but another error was produced")
		}
		log.Fatal(err)
	}
	event.TriggerEvent(event.LocalConfigComplete)
}

func GetAppConfigName(name, active string) (string, error) {
	if name == "" {
		return "", errors.New("cc.application.name not config")
	} else if active == "" {
		return "", errors.New("cc.profiles.active not config")
	}
	return name + "-" + active + "." + LocalNacosConfig.FileExtension, nil
}

func GetStringOrDefault(key, defaultVal string) string {
	str := viper.GetString(key)
	if str == "" {
		return defaultVal
	}
	return str
}

func Default(key string, value interface{}) {
	viper.SetDefault(key, value)
}

func Sub(key string) *viper.Viper { return viper.Sub(key) }

func GetString(key string) string { return viper.GetString(key) }

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}
func GetInt64(key string) int64 { return viper.GetInt64(key) }
