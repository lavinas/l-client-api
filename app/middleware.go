// Package app is a tools set that respond for a http deal layer 
package app

import (
	"net/http"
	"time"

	"github.com/lavinas/l-client-api/util/logger"
)
// middleware is a function that start all middlewares
func middleware() {
	r.Use(logMiddleware)
}
// logResponseWriter is wrapper for take response writer attributes
type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}
// NewlogResponseWriter is a function that starts a new lofResponseWriter
func NewlogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}
// WriteHeader is a logResponseWriter wrapper method that breaks apart responsewroter attributes
func (l *logResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}
// logMiddleware is a function that returns a middleware responsible for logging http calls
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := NewlogResponseWriter(w)
		s := time.Now()
		next.ServeHTTP(l, r)
		el := time.Since(s)
		rl := logger.RequestLogger{Uri: r.RequestURI, Addr: r.RemoteAddr, Code: l.statusCode, Duration: el}
		rl.Request()
	})
}
