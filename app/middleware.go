package app

import (
	"net/http"
	"time"

	"github.com/lavinas/l-client-api/util/logger"
)

func middleware() {
	r.Use(logMiddleware)
}

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewlogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (l *logResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.ResponseWriter.WriteHeader(code)
}

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := NewlogResponseWriter(w)
		s := time.Now()
		next.ServeHTTP(l, r)
		el := time.Since(s)
		rl := logger.RequestLogger{Uri: r.RequestURI, Addr: r.RemoteAddr, Code: l.statusCode, Duration: el}
		rl.Info()
    })
}