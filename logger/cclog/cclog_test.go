package cclog_test

import (
	"github.com/cqu20141693/go-service-common/logger/cclog"
	"go.uber.org/zap/zapcore"
	"testing"
)

func init() {
	cclog.SetLevel(zapcore.DebugLevel)
}
func TestLog(t *testing.T) {

	cclog.Debug("debug")
	cclog.Info("info")
	cclog.Warn("warn")
	cclog.Error("error")
}
