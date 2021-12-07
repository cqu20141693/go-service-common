package event

import (
	"context"
	"github.com/cqu20141693/go-service-common/logger"
)

// event
type MicroEvent int8

const (
	Init MicroEvent = iota
	Start
	LocalConfigComplete
	ConfigComplete
	RouterRegisterComplete
)

type ConfigHook func(ctx context.Context)

var concurrent = Init

type HookContext struct {
	hook ConfigHook
	Ctx  context.Context
}

func NewHookContext(hook ConfigHook, ctx context.Context) *HookContext {
	return &HookContext{hook: hook, Ctx: ctx}
}

var HookMap = make(map[MicroEvent][]*HookContext)

func RegisterHook(e MicroEvent, hookCtx *HookContext) {
	if concurrent >= e {
		hookCtx.hook(hookCtx.Ctx)
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
	if concurrent < event {
		for _, hookCtx := range HookMap[event] {
			hookCtx.hook(hookCtx.Ctx)
		}
		concurrent = event
	} else {
		for _, hookCtx := range HookMap[event-1] {
			hookCtx.hook(hookCtx.Ctx)
		}
		concurrent = event
		logger.Info(context.Background(), "current event must be less event")
	}

}
