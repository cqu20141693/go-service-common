package boot

import (
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/web"
)

func Boot(args []string) {
	event.TriggerEvent(event.Start)
	event.RegisterHook(event.RouterRegisterComplete, web.Start)
}
