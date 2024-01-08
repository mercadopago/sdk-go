package api

import (
	"net/http"

	"github.com/mercadopago/sdk-go/pkg/httpclient"
)

type Options struct {
	Requester httpclient.Requester
}

type requestOptions struct {
	customHeaders http.Header

	Options
}

type Option interface {
	RequestOption
	ApplyOption(opts *Options)
}

type RequestOption interface {
	applyRequestOption(opts *requestOptions)
}

type optFunc func(opts *Options)

func (f optFunc) ApplyOption(o *Options)               { f(o) }
func (f optFunc) applyRequestOption(o *requestOptions) { f(&o.Options) }

type requestOptFunc func(opts *requestOptions)

func (f requestOptFunc) applyRequestOption(o *requestOptions) { f(o) }

// WithRequester allow do request api call using received requester.
func WithRequester(r httpclient.Requester) Option {
	return optFunc(func(options *Options) {
		if r != nil {
			options.Requester = r
		}
	})
}

// WithCustomHeaders set request headers, it will be send to api.
func WithCustomHeaders(h http.Header) RequestOption {
	return requestOptFunc(func(options *requestOptions) {
		options.customHeaders = h
	})
}
