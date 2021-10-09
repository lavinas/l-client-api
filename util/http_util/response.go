// Package http_util provides tools for standardize and help http interface actions
package http_util

import (
	"encoding/json"
	"net/http"
)

// RespondJson is a fuction that format and standardize http json response
func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

// RespondError is a fuction that format and standardize http json with a error
func RespondError(w http.ResponseWriter, err RestErr) {
	RespondJson(w, err.Status(), err)
}
