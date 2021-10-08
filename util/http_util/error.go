package http_util

import (
	"errors"
	"encoding/json"
	"fmt"
	"net/http"
)

type RestErr interface {
	Error() string
	Message() string
	Status() int
}

type restErr struct {
	Rmessage string `json:"message"`
	Rstatus  int    `json:"status"`
	Rerror   string `json:"error"`
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var rErr restErr
	if err := json.Unmarshal(bytes, &rErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return rErr, nil
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s", e.Rmessage, e.Rstatus, e.Rerror)
}

func (e restErr) Message() string {
	return e.Rmessage
}

func (e restErr) Status() int {
	return e.Rstatus
}

func NewRestErr(message string, status int, error string) RestErr {
	return restErr{message, status, error}
}

func NewBadRequestError(message string) RestErr {
	return NewRestErr(message, http.StatusBadRequest, "bad_request")
}

func NewNotFoundError(message string) RestErr {
	return NewRestErr(message, http.StatusNotFound, "not_found")
}

func NewInternalServerError(message string) RestErr {
	return NewRestErr(message, http.StatusInternalServerError, "internal_server_error")
}

func NewNotImplementedError(message string) RestErr {
	return NewRestErr(message, http.StatusNotImplemented, "not_implemented")
}

func NewUnauthorizedError(message string) RestErr {
	return NewRestErr(message, http.StatusUnauthorized, "unauthorized")
}
