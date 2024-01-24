package option

import "github.com/mercadopago/sdk-go/pkg/internal/httpclient"

type ClientOptions struct {
	Requester httpclient.Requester
}

// ClientOption signature for client configurable parameters.
type ClientOption interface {
	Apply(opts *ClientOptions)
}

type optFunc func(opts *ClientOptions)

func (f optFunc) Apply(o *ClientOptions) { f(o) }

// WithCustomClient allow do requests using received requester.
func WithCustomClient(r httpclient.Requester) ClientOption {
	return optFunc(func(options *ClientOptions) {
		if r != nil {
			options.Requester = r
		}
	})
}

func ApplyClientOptions(opts ...ClientOption) *ClientOptions {
	options := ClientOptions{
		Requester: httpclient.DefaultRequester(),
	}
	for _, opt := range opts {
		opt.Apply(&options)
	}

	return &ClientOptions{
		Requester: options.Requester,
	}
}
