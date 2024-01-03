package httpclient

import (
	"net/http"
)

type requestOptions struct {
	RequestRequester Requester
	CustomHeaders    http.Header
}

type RequestOption interface {
	applyRequest(*requestOptions)
}

type requestOptFunc func(opts *requestOptions)

func (f requestOptFunc) applyRequest(o *requestOptions) { f(o) }

func WithRequestRequester(r Requester) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		if r != nil {
			options.RequestRequester = r
		}
	})
}

func WithCustomHeaders(h http.Header) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		options.CustomHeaders = h
	})
}
