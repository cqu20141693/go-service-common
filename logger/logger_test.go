package logger_test

import (
	"bytes"
	"context"
	zapctx "github.com/cqu20141693/go-service-common/logger"
	"github.com/cqu20141693/go-service-common/logger/cclog"
	qt "github.com/frankban/quicktest"
	"github.com/juju/loggo"
	"github.com/nacos-group/nacos-sdk-go/common/file"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)
	ctx := zapctx.WithLogger(context.Background(), logger)
	zapctx.Logger(ctx).Info("hello")
	c.Assert(buf.String(), qt.Matches, `INFO\thello\n`)
}

func TestDefaultLogger(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)
	ctx := zapctx.WithLogger(context.Background(), logger)
	zapctx.Logger(ctx).Info("hello")
	c.Assert(buf.String(), qt.Matches, `INFO\thello\n`)
}

func TestWithFields(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)

	ctx := zapctx.WithLogger(context.Background(), logger)
	ctx = zapctx.WithFields(ctx, zap.Int("foo", 999), zap.String("bar", "whee"))
	zapctx.Logger(ctx).Info("hello")
	c.Assert(buf.String(), qt.Matches, `INFO\thello\t\{"foo": 999, "bar": "whee"\}\n`)
}

func TestWithLevel(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)

	ctx := zapctx.WithLogger(context.Background(), logger)
	ctx1 := zapctx.WithLevel(ctx, zap.WarnLevel)
	zapctx.Info(ctx, "one")
	zapctx.Info(ctx1, "should not appear")
	zapctx.Warn(ctx1, "two")
	zapctx.Error(ctx1, "three")
	c.Assert(buf.String(), qt.Matches, `INFO\tone\nWARN\ttwo\nERROR\tthree\n`)
}

func TestMultistageSetupA(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)

	ctx := zapctx.WithLogger(context.Background(), logger)
	ctx = zapctx.WithLevel(ctx, zapcore.WarnLevel)
	ctx = zapctx.WithFields(ctx, zap.String("foo", "bar"))
	zapctx.Info(ctx, "one")
	zapctx.Warn(ctx, "two")
	c.Assert(buf.String(), qt.Matches, `WARN\ttwo\t{\"foo\": \"bar\"}\n`)
}

func TestMultistageSetupB(t *testing.T) {
	c := qt.New(t)
	var buf bytes.Buffer
	logger := cclog.NewLogger(&buf)

	ctx := zapctx.WithLogger(context.Background(), logger)
	ctx = zapctx.WithFields(ctx, zap.String("foo", "bar"))
	ctx = zapctx.WithLevel(ctx, zapcore.WarnLevel)
	zapctx.Info(ctx, "one")
	zapctx.Warn(ctx, "two")
	c.Assert(buf.String(), qt.Matches, `WARN\ttwo\t{\"foo\": \"bar\"}\n`)
}

func TestZapUtil(t *testing.T) {
	c := qt.New(t)
	// buffer
	var buf bytes.Buffer
	entry := loggo.Entry{Message: "test-log"}
	testWriterBuffer(&buf, entry)
	c.Assert(buf.String(), qt.Matches, "INFO\ttest-log\t{\"module\": \"\", \"timestamp\": \"0001-01-01 00:00:00.000\"}\n")
	//console
	testWriterConsole(os.Stdout, entry)
	// rotatelogs
	service := "sip-service"
	rotateTime := "24h"
	maxAge := 3
	writer, err := cclog.GetWriter(file.GetCurrentPath(), service+".log", rotateTime, int64(maxAge))
	if err != nil {
		logger.Info("get log writer failed")
		return
	}
	testWriterRotate(writer, entry)
}

func testWriterRotate(w io.Writer, entry loggo.Entry) {
	entry.Timestamp = time.Now()
	logger := cclog.NewLogger(w)
	writer := cclog.NewLoggoWriter(logger)
	writer.Write(entry)
}

func testWriterBuffer(w io.Writer, entry loggo.Entry) {
	logger := cclog.NewLogger(w)
	writer := cclog.NewLoggoWriter(logger)
	writer.Write(entry)
}
func testWriterConsole(w io.Writer, entry loggo.Entry) {
	entry.Timestamp = time.Now()
	logger := cclog.NewLogger(w)
	writer := cclog.NewLoggoWriter(logger)
	writer.Write(entry)
}
