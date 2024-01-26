package httpclient

import "net/http"

// ResponseError represents an error response from the API.
type ResponseError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`

	Headers http.Header `json:"headers"`
}

// Error implements error.
func (e *ResponseError) Error() string {
	return e.Message
}
