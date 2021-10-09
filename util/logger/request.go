// Package logger provides logger tools
// logger use go.uber.org/zap as support main tool
package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
//  RequestLogger represents information of the request and response
//  that will be placed on log
type RequestLogger struct {
	// Uri has the api's uri that was called
	Uri string
	// Addr has the address of the caller
	Addr string
	// Code has the http result code
	Code int
	// Duration has the duration of the api called
	Duration time.Duration
}
// MarshalLogObject is a RequestLogger method that encapsule fields in a zapcore encoder
func (r *RequestLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("uri", r.Uri)
	enc.AddString("addr", r.Addr)
	enc.AddInt("code", r.Code)
	enc.AddString("dur", fmt.Sprintf("%s ", r.Duration))
	return nil
}
// Info is a RequestLogger method that register a request log line in logger adding RequestLogger info
func (r *RequestLogger) Request() {
	Info("request", zap.Object("request", r))
}
// Info is a RequestLogger method that register a log line in logger adding RequestLogger info
func (r *RequestLogger) Info(msg string) {
	Info(msg, zap.Object("request", r))
}
// Error is a RequestLogger method that register a log line in logger adding RequestLogger info
func (r *RequestLogger) Error(msg string, err error) {
	Error(msg, err, zap.Object("request", r))
}
// Fatal is a RequestLogger method that register a log line in logger adding RequestLogger info
func (r *RequestLogger) Fatal(msg string, err error) {
	Fatal(msg, err, zap.Object("request", r))
}
