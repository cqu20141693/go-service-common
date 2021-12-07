package boot

import (
	"context"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/web"
	signalutil "go-micro.dev/v4/util/signal"
	"os"
	"os/signal"
)

func Boot(args []string) {
	event.TriggerEvent(event.Start)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalutil.Shutdown()...)

	background := context.Background()
	event.RegisterHook(event.RouterRegisterComplete, event.NewHookContext(web.Start, background))
	select {
	// wait on kill signal
	case <-ch:
	}

}
