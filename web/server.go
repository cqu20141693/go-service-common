package web

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/v2/config"
	"github.com/cqu20141693/go-service-common/v2/event"
	"github.com/cqu20141693/go-service-common/v2/logger/cclog"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/util/log"
	"os"
)

func init() {

	event.RegisterHook(event.LogComplete, event.NewHookContext(InitRouter, "initRouter"))
}

var Engine = gin.New()

type RouterRegister func(router *gin.Engine)

var routerRegisterSlice = make([]RouterRegister, 0)

func AddRouterRegister(register RouterRegister) {
	if register != nil {
		routerRegisterSlice = append(routerRegisterSlice, register)
	}
	cclog.Debug(fmt.Sprintf("AddRouterRegister size=%d", len(routerRegisterSlice)))
}

func InitRouter() {
	for _, register := range routerRegisterSlice {
		register(Engine)
	}
	cclog.Debug(fmt.Sprintf("trigger RouterRegisterComplete size=%d", len(routerRegisterSlice)))
	event.TriggerEvent(event.RouterRegisterComplete)
}

func Start() {
	cclog.Debug("web start")
	go func() {
		address := config.GetStringOrDefault("server.port", "8080")
		err := Engine.Run(":" + address)
		if err != nil {
			log.Error(err)
			os.Exit(0)
			return
		}
	}()

}
