package option

type ClientOptions struct {
	Requester Requester
}

// ClientOption signature for client configurable parameters.
type ClientOption interface {
	Apply(opts *ClientOptions)
}

type optFunc func(opts *ClientOptions)

func (f optFunc) Apply(o *ClientOptions) { f(o) }

// WithCustomClient allow do requests using received requester.
func WithCustomClient(r Requester) ClientOption {
	return optFunc(func(options *ClientOptions) {
		if r != nil {
			options.Requester = r
		}
	})
}
