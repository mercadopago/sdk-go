package httpclient

type ApiOptions struct {
	APIRequester Requester
}

// APIOption signature for api client configurable parameters.
type APIOption func(*ApiOptions)

// WithAPIRequester allow do api client requests using received requester.
func WithAPIRequester(r Requester) APIOption {
	return func(options *ApiOptions) {
		if r != nil {
			options.APIRequester = r
		}
	}
}
