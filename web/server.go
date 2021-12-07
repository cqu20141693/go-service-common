package web

import (
	"context"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {

	event.RegisterHook(event.ConfigComplete, event.NewHookContext(InitRouter, context.Background()))
}

var Engine = gin.New()

type RouterRegister func(router *gin.Engine)

var routerRegisterSlice = make([]RouterRegister, 0)

func AddRouterRegister(register RouterRegister) {
	if register != nil {
		routerRegisterSlice = append(routerRegisterSlice, register)
	}
}

func InitRouter(ctx context.Context) {
	for _, register := range routerRegisterSlice {
		register(Engine)
	}
	event.TriggerEvent(event.RouterRegisterComplete)
}

func Start(ctx context.Context) {
	go func() {
		address := config.GetStringOrDefault("server.port", "8080")
		err := Engine.Run(":" + address)
		if err != nil {
			cclog.Error("api server start failed")
			os.Exit(0)
			return
		}
	}()

}
