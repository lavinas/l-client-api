// Package app is a tools set that respond for a http deal layer 
package app

import (
	"net/http"

	"github.com/lavinas/l-client-api/handler"
)
// maps is a function that map http handlers
func maps() {
	r.HandleFunc("/ping", handler.Ping.Ping).Methods(http.MethodGet)
}
