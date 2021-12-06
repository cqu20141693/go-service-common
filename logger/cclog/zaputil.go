package cclog

import (
	"github.com/juju/loggo"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/nacos-group/nacos-sdk-go/common/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"time"
)

var loggoToZap = map[loggo.Level]zapcore.Level{
	loggo.TRACE:    zap.DebugLevel, // There's no zap equivalent to TRACE.
	loggo.DEBUG:    zap.DebugLevel,
	loggo.INFO:     zap.InfoLevel,
	loggo.WARNING:  zap.WarnLevel,
	loggo.ERROR:    zap.ErrorLevel,
	loggo.CRITICAL: zap.ErrorLevel, // There's no zap equivalent to CRITICAL.
}

// NewLoggoWriter returns a loggo.Writer that writes to the
// given zap logger.
func NewLoggoWriter(logger *zap.Logger) loggo.Writer {
	return zapLoggoWriter{
		logger: logger,
	}
}

// zapLoggoWriter implements a loggo.Writer by writing to a zap.Logger,
// so can be used as an adaptor from loggo to zap.
type zapLoggoWriter struct {
	logger *zap.Logger
}

// zapLoggoWriter implements loggo.Writer.Write by writing the entry
// to w.logger. It ignores entry.Timestamp because zap will affix its
// own timestamp.
func (w zapLoggoWriter) Write(entry loggo.Entry) {
	if ce := w.logger.Check(loggoToZap[entry.Level], entry.Message); ce != nil {
		ce.Write(zap.String("module", entry.Module), zap.String("timestamp", entry.Timestamp.Format("2006-01-02 15:04:05.000")))
	}
}

func GetWriter(outputPath, name, rotateTime string, maxAge int64) (writer io.Writer, err error) {
	err = file.MkdirIfNecessary(outputPath)
	if err != nil {
		return
	}
	outputPath = outputPath + string(os.PathSeparator)
	rotateDuration, err := time.ParseDuration(rotateTime)
	writer, err = rotatelogs.New(filepath.Join(outputPath, name+"-%Y%m%d%H%M"),
		rotatelogs.WithRotationTime(rotateDuration), rotatelogs.WithMaxAge(time.Duration(maxAge)*rotateDuration),
		rotatelogs.WithLinkName(filepath.Join(outputPath, name)))
	return
}

var defaultLevel = zapcore.InfoLevel

func SetLevel(level zapcore.Level) {
	defaultLevel = level
}

func NewLogger(w io.Writer) *zap.Logger {
	config := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		EncodeLevel: zapcore.CapitalLevelEncoder,
	}
	level := zapcore.InfoLevel
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(w),
		level,
	)
	return zap.New(core)
}
