package app

import (
	"net/http"

	"github.com/lavinas/l-client-api/handler"
)

func maps() {
	r.HandleFunc("/ping", handler.Ping.Ping).Methods(http.MethodGet)
}