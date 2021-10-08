package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RequestLogger struct {
	Uri string
	Addr string
	Response int

}

func (r *RequestLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("uri", r.Uri)
	enc.AddString("addr", r.Addr)
	enc.AddInt("response", r.Response)
	return nil
}

func (r *RequestLogger) Info() {
	log.log.Info("request", zap.Object("request_info", r))
	log.log.Sync()
}