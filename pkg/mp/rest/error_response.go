package rest

import "net/http"

// ErrorResponse represents an error response from the API.
type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`

	Headers http.Header `json:"headers"`
}

// Error implements error.
func (e *ErrorResponse) Error() string {
	return e.Message
}
