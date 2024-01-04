package httpclient

type ApiOptions struct {
	APIRequester Requester
}

// APIOption signature for api client configurable parameters.
type APIOption interface {
	applyAPI(*ApiOptions)
}

type apiOptFunc func(opts *ApiOptions)

func (f apiOptFunc) applyAPI(o *ApiOptions) { f(o) }

// WithAPIRequester allow do api client requests using received requester.
func WithAPIRequester(r Requester) APIOption {
	return apiOptFunc(func(options *ApiOptions) {
		if r != nil {
			options.APIRequester = r
		}
	})
}

func BuildAPIClientOptions(opts ...APIOption) *ApiOptions {
	options := ApiOptions{
		APIRequester: NewRetryable(3),
	}

	for _, opt := range opts {
		opt.applyAPI(&options)
	}

	return &options
}
