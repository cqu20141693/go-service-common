package cclog_test

import (
	ccboot "github.com/cqu20141693/go-service-common/v2/boot"
	"github.com/cqu20141693/go-service-common/v2/global"
	"github.com/cqu20141693/go-service-common/v2/logger/cclog"
	"go.uber.org/zap/zapcore"
	"testing"
)

func init() {
	cclog.SetLevel(zapcore.DebugLevel)
	global.SetLogLevel(zapcore.DebugLevel)
}
func TestLog(t *testing.T) {
	ccboot.Boot(nil)
	cclog.Debug("debug")
	cclog.Info("info")
	cclog.Warn("warn")
	cclog.Error("error")
}
