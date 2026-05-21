// Package mperror defines error types returned by the MercadoPago Go SDK.
//
// When the API responds with an HTTP status code >= 400, the SDK wraps the
// failure in a [ResponseError] that carries the status code, response headers,
// and the raw error message from the API body. Callers can use a type assertion
// or errors.As to inspect the structured error:
//
//	var mpErr *mperror.ResponseError
//	if errors.As(err, &mpErr) {
//	    log.Printf("API error %d: %s", mpErr.StatusCode, mpErr.Message)
//	}
package mperror

import "net/http"

// ResponseError represents a non-success HTTP response from the MercadoPago API.
// It implements the built-in error interface so it can be returned and inspected
// through standard Go error-handling patterns.
type ResponseError struct {
	// Headers contains the HTTP response headers returned by the API. Useful for
	// debugging rate-limit or correlation headers.
	Headers http.Header `json:"headers"`

	// Message holds the raw error message extracted from the API response body.
	Message string `json:"message"`

	// StatusCode is the HTTP status code of the failed response (e.g. 400, 404, 500).
	StatusCode int `json:"status_code"`
}

// Error returns the API error message, satisfying the built-in error interface.
func (e *ResponseError) Error() string {
	return e.Message
}
