package httpclient

import (
	"net/http"
)

type requestOptions struct {
	callRequester Requester
	customHeaders http.Header
}

// RequestOption signature for client request configurable parameters.
type RequestOption interface {
	applyRequest(*requestOptions)
}

type requestOptFunc func(opts *requestOptions)

func (f requestOptFunc) applyRequest(o *requestOptions) { f(o) }

// WithCallRequester allow do request api call using received requester.
func WithCallRequester(r Requester) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		if r != nil {
			options.callRequester = r
		}
	})
}

// WithCustomHeaders set request headers, it will be send to api.
func WithCustomHeaders(h http.Header) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		options.customHeaders = h
	})
}
