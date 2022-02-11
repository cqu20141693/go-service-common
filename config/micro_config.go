package config

import (
	"bytes"
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"github.com/asim/go-micro/plugins/config/source/nacos/v4"
	"github.com/cqu20141693/go-service-common/v2/event"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/spf13/viper"
	config4 "go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source"
	"go-micro.dev/v4/config/source/file"
	"log"
)

var encoder = yaml.NewEncoder()
var Conf config4.Config
var err error
var Viper = viper.New()

func InitV2() {

	Conf, err = config4.NewConfig(config4.WithReader(json.NewReader(reader.WithEncoder(encoder))))
	if err != nil {
		log.Fatal(err)
	}
	event.RegisterHook(event.Start, event.NewHookContext(readConfigV2, "initConfig"))
}

func readConfigV2() {
	loadLocalFile("./resource/bootstrap.yaml", Conf)

	active := Conf.Get("cc", "profiles", "active").String("")
	if active != "" {
		loadLocalFile("./resource/bootstrap-"+active+".yaml", Conf)
	}
	event.TriggerEvent(event.LocalConfigComplete)
	loadNacosConfig(Conf, active)
	Viper.SetConfigType("json")
	err := Viper.MergeConfig(bytes.NewBuffer(Conf.Bytes()))
	if err != nil {
		log.Fatal("Merge Config failed")
	}
	event.TriggerEvent(event.ConfigComplete)
}

func loadNacosConfig(conf config4.Config, active string) {
	appName := conf.Get("cc", "application", "name").String("")

	if appName != "" {
		addr := conf.Get("cc", "cloud", "nacos", "address").StringSlice(nil)
		if addr == nil {
			return
		}
		clientConfig := getClientConfig(conf)
		loadAppConfig(addr, clientConfig, getDataId(conf, appName, active), getGroup(conf), conf)
		fmt.Println("nacosAppMap=", conf.Map())
		loadExtConfig(conf, addr, clientConfig)

	}
}
func loadAppConfig(addr []string, cc constant.ClientConfig, dataId, group string, conf config4.Config) {
	var ops []source.Option
	ops = append(ops, nacos.WithAddress(addr))
	ops = append(ops, nacos.WithClientConfig(cc))
	ops = append(ops, nacos.WithDataId(dataId))
	ops = append(ops, nacos.WithGroup(group))
	ops = append(ops, source.WithEncoder(encoder))
	err = conf.Load(nacos.NewSource(ops...))
	if err != nil {
		log.Printf("load nacos config failed,dataId=%s err=%s\n", dataId, err.Error())
	}
}

func getGroup(conf config4.Config) string {

	return conf.Get("cc", "cloud", "nacos", "group").String("DEFAULT_GROUP")
}

func getDataId(conf config4.Config, appName string, active string) string {
	fileExt := conf.Get("cc", "cloud", "nacos", "file-extension").String("yaml")
	dataId := appName + "-" + active + "." + fileExt
	return dataId
}

func getClientConfig(conf config4.Config) constant.ClientConfig {
	clientConfig := constant.ClientConfig{}
	configValue := conf.Get("cc", "cloud", "nacos", "config")
	if configValue.StringMap(nil) != nil {
		_ = configValue.Scan(&clientConfig)
	}
	return clientConfig
}

func loadExtConfig(conf config4.Config, addr []string, clientConfig constant.ClientConfig) {
	configs := conf.Get("cc", "cloud", "nacos", "extend-configs")
	type ExtConfig struct {
		DataId  string
		Group   string
		Refresh bool
	}
	var extConfigs []ExtConfig
	err := configs.Scan(&extConfigs)
	if err != nil {
		log.Println(err)
	}
	if extConfigs != nil && len(extConfigs) > 0 {
		for _, extConfig := range extConfigs {
			if extConfig.Group != "" {
				loadAppConfig(addr, clientConfig, extConfig.DataId, extConfig.Group, conf)
			} else {
				loadAppConfig(addr, clientConfig, extConfig.DataId, "DEFAULT_GROUP", conf)
			}
		}
	}
}
func loadLocalFile(path string, conf config4.Config) {
	fileSource := file.NewSource(file.WithPath(path), source.WithEncoder(yaml.NewEncoder()))
	err = conf.Load(fileSource)
	if err != nil {
		log.Println("load file failed", err)
		return
	}
}
