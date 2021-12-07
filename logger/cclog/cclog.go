package cclog

import (
	"context"
	"github.com/cqu20141693/go-service-common/config"
	"github.com/cqu20141693/go-service-common/event"
	"github.com/cqu20141693/go-service-common/file"
	"github.com/juju/loggo"
	"os"
	"strings"
	"time"
)

var logs []loggo.Writer
var console loggo.Writer

func init() {
	event.RegisterHook(event.ConfigComplete, event.NewHookContext(configRotate, context.Background()))
}

func configRotate(ctx context.Context) {

	stdout := NewLogger(os.Stdout)
	console = NewLoggoWriter(stdout)
	logs = []loggo.Writer{console}
	addDefaultConfig()
	rotateTime := config.GetStringOrDefault("cc.log.rotate-time", "24h")
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
	writer, err := GetWriter(path, service+".log", rotateTime, maxAge)
	if err != nil {
		console.Write(ErrorEntry("rotate writer create failed"))
		return
	}
	rotate := NewLoggoWriter(NewLogger(writer))
	logs = append(logs, rotate)
}

func addDefaultConfig() {
	config.Default("cc.log.max-age", 3)
}

func DebugEntry(msg string) loggo.Entry {
	return loggo.Entry{Level: loggo.DEBUG, Message: msg, Timestamp: time.Now()}
}
func InfoEntry(msg string) loggo.Entry {
	return loggo.Entry{Level: loggo.INFO, Message: msg, Timestamp: time.Now()}
}
func WarnEntry(msg string) loggo.Entry {
	return loggo.Entry{Level: loggo.WARNING, Message: msg, Timestamp: time.Now()}
}
func ErrorEntry(msg string) loggo.Entry {
	return loggo.Entry{Level: loggo.ERROR, Message: msg, Timestamp: time.Now()}
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
