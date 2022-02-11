package global

import (
	"github.com/cqu20141693/go-service-common/v2/logger/cclog"
	"go.uber.org/zap/zapcore"
	"os"
)

var defaultLogLevel = zapcore.InfoLevel

func GetLogLevel() zapcore.Level {
	return defaultLogLevel
}
func SetLogLevel(level zapcore.Level) {
	defaultLogLevel = level
	cclog.AddLogger("console", zapcore.DebugLevel, os.Stdout)
}
