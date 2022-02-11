package boot

import (
	"encoding/json"
	"fmt"
	"github.com/cqu20141693/go-service-common/v2/event"
	"github.com/cqu20141693/go-service-common/v2/logger/cclog"
	"github.com/cqu20141693/go-service-common/v2/web"
)

func Boot(args []string) {

	event.TriggerEvent(event.Start)
	marshal, _ := json.Marshal(args)
	cclog.Debug(fmt.Sprintf("args=%s", string(marshal)))

	event.RegisterHook(event.RouterRegisterComplete, event.NewHookContext(web.Start, "webStart"))

	ListenSignal()

}
