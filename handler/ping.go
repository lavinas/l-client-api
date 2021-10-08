package handler

import (
	"net/http"

	"github.com/lavinas/l-client-api/util/http_util"
)

const (
	pong = "pong"
)

var (
	Ping pingHandlerInterface = &pingHandler{}
)

type pingHandlerInterface interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingHandler struct{}

func (c *pingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	http_util.RespondJson(w, http.StatusOK, pong)
}
