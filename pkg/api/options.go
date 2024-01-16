package api

import (
	"net/http"
)

type requestOptions struct {
	httpClient    *http.Client
	customHeaders http.Header
}

// RequestOption signature for client configurable parameters by request.
type RequestOption interface {
	applyRequestOption(opts *requestOptions)
}

type requestOptFunc func(opts *requestOptions)

func (f requestOptFunc) applyRequestOption(o *requestOptions) { f(o) }

// WithHTTPClient allow do request api call using received http client.
func WithHTTPClient(r *http.Client) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		if r != nil {
			options.httpClient = r
		}
	})
}

// WithCustomHeaders set request headers, it will be send to api.
func WithCustomHeaders(h http.Header) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		options.customHeaders = h
	})
}
