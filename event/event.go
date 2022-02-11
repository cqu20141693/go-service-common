package event

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/v2/logger/cclog"
)

// event
type MicroEvent int8

const (
	Init MicroEvent = iota
	Start
	LocalConfigComplete
	ConfigComplete
	LogComplete
	RouterRegisterComplete
)

type ConfigHook func()

var concurrent = Init

type HookContext struct {
	hook ConfigHook
	name string
}

func NewHookContext(hook ConfigHook, name string) *HookContext {
	return &HookContext{hook: hook, name: name}
}

var HookMap = make(map[MicroEvent][]*HookContext)

func RegisterHook(e MicroEvent, hookCtx *HookContext) {
	cclog.Debug(fmt.Sprintf("RegisterHook [%s] on event=%d", hookCtx.name, e))
	if concurrent >= e {
		hookCtx.hook()
	} else {
		hooks, ok := HookMap[e]
		if !ok {
			hooks = make([]*HookContext, 0)
		}
		hooks = append(hooks, hookCtx)
		HookMap[e] = hooks
	}
}

func TriggerEvent(event MicroEvent) {
	cclog.Debug(fmt.Sprintf("trigger event=%d", event))
	if concurrent < event {
		concurrent = event
		for _, hookCtx := range HookMap[event] {
			hookCtx.hook()
		}
	} else {
		cclog.Info("current event must be less event")
	}

}
