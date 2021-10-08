package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type loggerInterface interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

type logger struct {
	log *zap.Logger
}

func GetLogger() loggerInterface {
	return log
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
		return
	}
	Info(fmt.Sprintf(format, v...))
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v...))
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutput()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func getOutput() string {
	ret := os.Getenv(strings.ToLower(strings.TrimSpace(envLogOutput)))
	if ret == "" {
		return "output"
	}
	return ret
}

func getLevel() zapcore.Level {
	switch os.Getenv(strings.ToLower(strings.TrimSpace(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}

func Fatal(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Fatal(msg, tags...)
}