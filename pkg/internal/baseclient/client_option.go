package baseclient

// ClientOption allows sending options in the http baseclient
type clientOption struct {
	pathParams  map[string]string
	queryParams map[string]string
}

type Option func(*clientOption)

// WithPathParams allows sending path parameters in the request.
func WithPathParams(params map[string]string) Option {
	return func(c *clientOption) {
		if params != nil {
			c.pathParams = params
		}
	}
}

// WithQueryParams allows sending query parameters in the request.
func WithQueryParams(params map[string]string) Option {
	return func(c *clientOption) {
		if params != nil {
			c.queryParams = params
		}
	}
}
