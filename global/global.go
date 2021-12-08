package global

import "go.uber.org/zap/zapcore"

var defaultLogLevel = zapcore.InfoLevel

func GetLogLevel() zapcore.Level {
	return defaultLogLevel
}
func SetLogLevel(level zapcore.Level) {
	defaultLogLevel = level
}
