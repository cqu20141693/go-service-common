package container

import (
	"context"
	"encoding/json"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/spf13/viper"
)

type ConfigProperties interface {
	prefix() string
}

var ConfigContainer []ConfigProperties = make([]ConfigProperties, 8)
var stop = make(chan bool)

func AddConfigProperties(cpc <-chan ConfigProperties) {
	for {
		select {

		case cp := <-cpc:
			ConfigContainer = append(ConfigContainer, cp)
			marshal, err := json.Marshal(viper.GetString(cp.prefix()))
			if err != nil {
				logger.Info(context.Background(), "config Marshal  failed")
				return
			}
			err = json.Unmarshal(marshal, &cp)
			if err != nil {
				logger.Info(context.Background(), "config Unmarshal failed")
				return
			}
		case <-stop:
			logger.Info(context.Background(), "ConfigProperties add stop")
			break
		}
	}
}

func ConfigUpdate() {
	logger.Info(context.Background(), "config update")
}
func Stop() {
	stop <- true
}
