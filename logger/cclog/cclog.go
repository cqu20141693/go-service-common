package cclog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var logs = map[string]Writer{"console": getConsoleWriter()}
var console Writer

func getConsoleWriter() Writer {
	stdout := newLogger(os.Stdout, zapcore.InfoLevel)
	console = NewCCLogWriter(stdout)
	return console
}
func newLogger(w io.Writer, level zapcore.Level) *zap.Logger {
	eConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalLevelEncoder,
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(eConfig),
		zapcore.AddSync(w),
		level,
	)
	return zap.New(core)
}

func AddWriter(key string, writer Writer) {
	logs[key] = writer
}
func AddLogger(key string, level zapcore.Level, writer io.Writer) {
	logs[key] = NewCCLogWriter(newLogger(writer, level))
}

func DebugEntry(msg string) Entry {
	return Entry{Level: DEBUG, Message: msg, Timestamp: time.Now()}
}
func InfoEntry(msg string) Entry {
	return Entry{Level: INFO, Message: msg, Timestamp: time.Now()}
}
func WarnEntry(msg string) Entry {
	return Entry{Level: WARNING, Message: msg, Timestamp: time.Now()}
}
func ErrorEntry(msg string) Entry {
	return Entry{Level: ERROR, Message: msg, Timestamp: time.Now()}
}

func Info(msg string) {
	for _, log := range logs {
		log.Write(InfoEntry(msg))
	}
}
func Debug(msg string) {
	for _, log := range logs {
		log.Write(DebugEntry(msg))
	}
}
func Warn(msg string) {
	for _, log := range logs {
		log.Write(WarnEntry(msg))
	}
}
func Error(msg string) {
	for _, log := range logs {
		log.Write(ErrorEntry(msg))
	}
}
