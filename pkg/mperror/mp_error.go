package mperror

import "net/http"

// ResponseError represents an error response from the API.
type ResponseError struct {
	Headers http.Header `json:"headers"`

	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Error implements error.
func (e *ResponseError) Error() string {
	return e.Message
}
