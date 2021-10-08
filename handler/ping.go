package handler

import (
	"net/http"
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
	w.Write([]byte(pong))
}
