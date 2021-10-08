package app

import (
	"encoding/json"
	"net/http"

	"github.com/lavinas/l-client-api/util/logger"
)

type logStruct struct {
	Uri string  `json:"uri"`
	Addr string `json:"addr"`
}

func middleware() {
	r.Use(logMiddleware)
}

func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logStruct{Uri: r.RequestURI, Addr: r.RemoteAddr}
		lj, _ := json.Marshal(l)
		logger.Info(string(lj))
        next.ServeHTTP(w, r)
    })
}