package config

import (
	"context"
	"errors"
	ccMicro "github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/spf13/viper"
	"log"
)

func init() {
	ReadLocalConfig()
	if viper.GetStringMap("cc.cloud.nacos.config") != nil {
		NacosInit()
	}
	ccMicro.TriggerEvent(ccMicro.ConfigComplete)
}
func ReadLocalConfig() {
	// 读取本地配置
	viper.SetConfigName("bootstrap.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")
	viper.AddConfigPath("/etc/resource")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info(context.Background(),"Config file not found; ignore error if desired")

		} else {
			logger.Info(context.Background(),"Config file was found but another error was produced")
		}
		log.Fatal(err)
	}
	ccMicro.TriggerEvent(ccMicro.LocalConfigComplete)
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
