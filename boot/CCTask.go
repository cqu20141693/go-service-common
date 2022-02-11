package boot

import (
	"github.com/cqu20141693/go-service-common/v2/config"
	"github.com/cqu20141693/go-service-common/v2/event"
	"github.com/cqu20141693/go-service-common/v2/logger"
	"github.com/cqu20141693/go-service-common/v2/utils"
	"os"
	"os/signal"
)

func init() {
	logger.Init()
	config.InitV2()
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
