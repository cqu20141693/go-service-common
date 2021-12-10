package boot

import (
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
)

func init() {
	logger.Init()
	config.Init()
}
func Task() {
	event.TriggerEvent(event.Start)
}
