package logger

import (
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"github.com/nacos-group/nacos-sdk-go/common/file"
	"os"
	"strings"
)

func Init() {
	event.RegisterHook(event.ConfigComplete, event.NewHookContext(ConfigRotate, "ConfigRotate"))
}
func ConfigRotate() {
	cclog.AddWriter("console", cclog.NewCCLogWriter(cclog.NewLogger(os.Stdout)))
	addDefaultConfig()
	rotateTime := config.GetString("cc.log.rotate-time")
	maxAge := config.GetInt64("cc.log.max-age")
	var path string
	if logDir := config.GetStringOrDefault("cc.log.dir", ""); logDir != "" {
		if strings.Contains(logDir, "/") {
			path = logDir
		} else {
			path = file.GetCurrentPath() + string(os.PathSeparator) + "log"
		}
	} else {
		path = file.GetCurrentPath()
	}

	service := config.GetStringOrDefault("cc.application.name", "service")
	writer, err := cclog.GetWriter(path, service+".log", rotateTime, maxAge)
	if err != nil {
		cclog.Error("rotate writer create failed")
		return
	}
	rotate := cclog.NewCCLogWriter(cclog.NewLogger(writer))
	cclog.AddWriter("rotate", rotate)
	cclog.Debug("trigger log complete")
	event.TriggerEvent(event.LogComplete)
}
func addDefaultConfig() {
	config.Default("cc.log.max-age", 3)
	config.Default("cc.log.rotate-time", "24h")
}
