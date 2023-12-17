// gofr_errors.go
package main

import (
	"encoding/json"
	"net/http"
)

// HTTPError represents an HTTP error response.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewHTTPError creates a new HTTPError with the given status code and message.
func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
	}
}

// RenderJSON renders the HTTPError as JSON.
func (e *HTTPError) RenderJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)

	json.NewEncoder(w).Encode(e)
}
