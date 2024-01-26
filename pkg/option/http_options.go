package option

import (
	"net/http"
)

// Requester has the minimum required method to send http requests.
type Requester interface {
	Do(req *http.Request) (*http.Response, error)
}

type ConfigOptions struct {
	HTTPClient Requester
}

// ConfigOption signature for client configurable parameters.
type ConfigOption interface {
	Apply(opts *ConfigOptions)
}

type optFunc func(opts *ConfigOptions)

func (f optFunc) Apply(o *ConfigOptions) { f(o) }

// WithCustomClient allow do requests using received http client.
func WithCustomClient(r Requester) ConfigOption {
	return optFunc(func(options *ConfigOptions) {
		if r != nil {
			options.HTTPClient = r
		}
	})
}
