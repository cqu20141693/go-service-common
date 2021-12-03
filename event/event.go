package event

import (
	"context"
	"github.com/cqu20141693/go-service-common/logger"
)

// event
type MicroEvent int8

const (
	Start MicroEvent = iota
	LocalConfigComplete
	ConfigComplete
)

type ConfigHook func()

var concurrent = Start
var HookMap = make(map[MicroEvent][]ConfigHook)

func RegisterHook(e MicroEvent, hook ConfigHook) {
	if concurrent >= e {
		hook()
	} else {
		hooks, ok := HookMap[e]
		if !ok {
			hooks = make([]ConfigHook, 0)
		}
		hooks = append(hooks, hook)
		HookMap[e] = hooks
	}
}

func TriggerEvent(event MicroEvent) {
	if concurrent < event {
		for _, hook := range HookMap[event] {
			hook()
		}
		concurrent = event
	} else {
		logger.Info(context.Background(),"current event must be less event")
	}

}
