package boot

import (
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-service-common/utils"
	"os"
	"os/signal"
)

func init() {
	logger.Init()
	config.Init()
}
func Task() {
	event.TriggerEvent(event.Start)
}

func ListenSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, utils.ShutDownSignals()...)

	// wait on kill signal
	select {

	case <-ch:
	}

}
