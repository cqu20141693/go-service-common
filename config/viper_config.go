package config

import (
	"errors"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Init() {
	event.RegisterHook(event.Start, event.NewHookContext(readConfig, "initConfig"))
}

func readConfig() {
	ReadLocalConfig()
	ReadCommandLine()
	ReadLocalProfile()
	nacosConfig := viper.GetStringMap("cc.cloud.nacos.config")
	if nacosConfig != nil && len(nacosConfig) > 0 {
		NacosInit()
	}
	event.TriggerEvent(event.ConfigComplete)
}

func ReadLocalProfile() {
	active := viper.GetString("cc.profiles.active")
	name := "bootstrap"
	if active != "" {
		fileName := name + "-" + active + ".yml"
		profileViper := viper.GetViper()
		profileViper.SetConfigName(fileName)
		profileViper.SetConfigType("yaml")
		profileViper.AddConfigPath("./resource")
		profileViper.AddConfigPath("/etc/resource")
		if err := profileViper.MergeInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				cclog.Debug("Config file not found; ignore error if desired")
			} else {
				cclog.Info("Config file was found but another error was produced")
			}
		}

	}
}

func ReadCommandLine() {
	pflag.String("cc.profiles.active", "", "config env variable")
	pflag.Parse()
	// 绑定命令行
	_ = viper.BindPFlags(pflag.CommandLine)
}
func ReadLocalConfig() {
	// 读取本地配置
	viper.SetConfigName("bootstrap.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resource")
	viper.AddConfigPath("/etc/resource")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cclog.Debug("Config file not found; ignore error if desired")
		} else {
			cclog.Info("Config file was found but another error was produced")
		}
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
