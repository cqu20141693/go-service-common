package boot

import (
	"encoding/json"
	"fmt"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/cqu20141693/go-service-common/web"
	signalutil "go-micro.dev/v4/util/signal"
	"os"
	"os/signal"
)

func init() {
	logger.Init()
	config.Init()
}
func Boot(args []string) {

	event.TriggerEvent(event.Start)
	marshal, _ := json.Marshal(args)
	cclog.Debug(fmt.Sprintf("args=%s", string(marshal)))
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signalutil.Shutdown()...)

	event.RegisterHook(event.RouterRegisterComplete, event.NewHookContext(web.Start, "webStart"))
	select {
	// wait on kill signal
	case <-ch:
	}

}
