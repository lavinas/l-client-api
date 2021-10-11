// Package logger provides logger tools
// logger uses go.uber.org/zap as support main tool
package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// envLogLevel store the log level enviroment name
	envLogLevel = "LOG_LEVEL"
	// envLogLevel store the log output enviroment name
	envLogOutput = "LOG_OUTPUT"
	// envLogErrorOutput store the log output error enviroment name
	envLogErrorOutput = "LOG_ERROR"
)

var (
	// log is a local static logger variable
	log logger
)

// loggerInterface represents a interface logger
type loggerInterface interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

// logger represents a internal logger struct
type logger struct {
	log *zap.Logger
}

// GetLogger returns a logger as a generic Logger interface
func GetLogger() loggerInterface {
	return log
}

// Printf is a logger method that implements a Prinft method for this logger
func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
		return
	}
	Info(fmt.Sprintf(format, v...))
}

// Printfis a logger method that implements a Prinf method for this logger
func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v...))
}

// init initialize module creating a zap logger as a logger variable
func init() {
	logConfig := zap.Config{
		OutputPaths:      []string{getOutputPath()},
		ErrorOutputPaths: []string{getErrorOutputPath()},
		Level:            zap.NewAtomicLevelAt(getLevel()),
		Encoding:         "json",
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

// getOutputPath returns the value of output path for zap logger initialization
// the path will be the path and file to store the log
// it try to capture from environment variable and will be the stdout as default
func getOutputPath() string {
	ret := os.Getenv(strings.ToLower(strings.TrimSpace(envLogOutput)))
	if ret == "" {
		return "stdout"
	}
	return ret
}

// getErrorOutputPath returns the value of output error path for zap logger initialization
// the path will be the path and file to store the log error
// it try to capture from environment variable and will be the stdout as default
func getErrorOutputPath() string {
	ret := os.Getenv(strings.ToLower(strings.TrimSpace(envLogErrorOutput)))
	if ret == "" {
		return "stdout"
	}
	return ret
}

// getLeval returns the value of level of logger
// it can be debug, info, error, panic or fatal
// it try to capture from environment variable and will be the info as default
func getLevel() zapcore.Level {
	switch os.Getenv(strings.ToLower(strings.TrimSpace(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "Fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// Info is a global function that store a debug level log
func Debug(msg string, tags ...zap.Field) {
	log.log.Debug(msg, tags...)
	log.log.Sync()
}

// Info is a global function that store a info level log
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

// Error is a global function that store a error level log
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}

// Info is a global function that store a panic level log
func Panic(msg string, tags ...zap.Field) {
	log.log.Panic(msg, tags...)
	log.log.Sync()
}

// Error is a global function that store a fatal level log
func Fatal(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Fatal(msg, tags...)
}
