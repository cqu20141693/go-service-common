package container

import (
	"context"
	"github.com/cqu20141693/go-service-common/logger"
	"syscall"
)

var SingletonFactory map[string]interface{}
var cpchan = make(chan ConfigProperties)

func init() {
	go AddConfigProperties(cpchan)
}
func InjectSingleton(key string, o interface{}) {
	if _, ok := SingletonFactory[key]; ok {
		logger.Info(context.Background(), "The singleton factory already exists instance="+key)
		syscall.Exit(1)
	}
	SingletonFactory[key] = o

	if cp, ok := o.(ConfigProperties); ok {
		cpchan <- cp
	}
}
