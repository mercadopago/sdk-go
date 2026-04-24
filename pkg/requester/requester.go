// Package requester defines the HTTP transport abstraction used by the
// MercadoPago Go SDK. By programming against the [Requester] interface, callers
// can substitute the default retry-capable transport with custom middleware,
// logging wrappers, or test mocks.
package requester

import "net/http"

// Requester is the minimal interface required to execute HTTP requests on behalf
// of the SDK. The default implementation (in the internal defaultrequester
// package) adds automatic retries with constant back-off, but any type whose Do
// method matches the standard [http.Client.Do] signature can be used instead.
type Requester interface {
	// Do sends an HTTP request and returns the corresponding HTTP response.
	// Implementations must close the request body (if any) and are free to add
	// retry logic, circuit-breaking, or observability as needed.
	Do(req *http.Request) (*http.Response, error)
}
