package app

import (
	// "fmt"
	"net/http"

	"github.com/lavinas/l-client-api/util/logger"
)

func middleware() {
	r.Use(logMiddleware)
}

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl := logger.RequestLogger{Uri: r.RequestURI, Addr: r.RemoteAddr}
		rl.Info()
        next.ServeHTTP(w, r)
    })
}