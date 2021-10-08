package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RequestLogger struct {
	Uri string
	Addr string
	Code int
	Duration time.Duration
}

func (r *RequestLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("uri", r.Uri)
	enc.AddString("addr", r.Addr)
	enc.AddInt("code", r.Code)
	enc.AddString("duration", fmt.Sprintf("%s ", r.Duration))
	return nil
}

func (r *RequestLogger) Info() {
	log.log.Info("request", zap.Object("request_info", r))
	log.log.Sync()
}