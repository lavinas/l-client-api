// Package http_util provides tools for standardize and help http interface actions
package http_util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)
// RestErr is a interface that standardize rest error tools
type RestErr interface {
	Error() string
	Message() string
	Status() int
}
// restErr is a struct who has a rest error variables
type restErr struct {
	// Rmessage has message error
	Rmessage string `json:"message"`
	// Rstatus has the http status code error number
	Rstatus  int    `json:"status"`
	// Rerror has the http status code error description
	Rerror   string `json:"error"`
}
// NewRestErrorFromBytes is a function transforms a byte string json to a RestError 
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var rErr restErr
	if err := json.Unmarshal(bytes, &rErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return rErr, nil
}
// Error is a RestError method that prints a error string
func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s", e.Rmessage, e.Rstatus, e.Rerror)
}
// Messsage is a RestError method that returns the message from RestError object
func (e restErr) Message() string {
	return e.Rmessage
}
// Messsage is a RestError method that returns the status code from RestError object
func (e restErr) Status() int {
	return e.Rstatus
}
// NewRestErr creates a new RestError object
func NewRestErr(message string, status int, error string) RestErr {
	return restErr{message, status, error}
}
// NewBadRequestError creates a object for a Bad Request Error
func NewBadRequestError(message string) RestErr {
	return NewRestErr(message, http.StatusBadRequest, "bad_request")
}
// NewNotFoundError creates a object for a Not Found Error
func NewNotFoundError(message string) RestErr {
	return NewRestErr(message, http.StatusNotFound, "not_found")
}
// NewInternalServerError creates a object for a Internal Server Error
func NewInternalServerError(message string) RestErr {
	return NewRestErr(message, http.StatusInternalServerError, "internal_server_error")
}
// NewNotImplementedError creates a object for a Not Implemented Error
func NewNotImplementedError(message string) RestErr {
	return NewRestErr(message, http.StatusNotImplemented, "not_implemented")
}
// NewUnauthorizedError creates a object for an Unauthorized Error
func NewUnauthorizedError(message string) RestErr {
	return NewRestErr(message, http.StatusUnauthorized, "unauthorized")
}
