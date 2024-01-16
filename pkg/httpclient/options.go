package httpclient

import (
	"net/http"
	"time"
)

type Options struct {
	RetryMax int

	HTTPClient      *http.Client
	Timeout         time.Duration
	BackoffStrategy BackoffFunc
	CheckRetry      CheckRetryFunc
}

// Option signature for client configurable parameters.
type Option interface {
	Apply(opts *Options)
}

type optFunc func(opts *Options)

func (f optFunc) Apply(o *Options) { f(o) }

// WithRetryMax tells the client the maximum number of retries to execute. Eg.: A
// value of 3, means to execute the original request, and up-to 3 retries (4
// requests in total). A value of 0 means no retries.
func WithRetryMax(max int) Option {
	return optFunc(func(options *Options) {
		options.RetryMax = max
	})
}

// WithHTTPClient allow do requests using received http client.
func WithHTTPClient(h *http.Client) Option {
	return optFunc(func(options *Options) {
		if h != nil {
			options.HTTPClient = h
		}
	})
}

// WithTimeout controls the timeout for each request. When retrying requests,
// each retried request will start counting from the beginning towards this
// timeout.
//
// A timeout of 0 disables request timeouts.
func WithTimeout(t time.Duration) Option {
	return optFunc(func(options *Options) {
		// Negative durations do not make sense in the context of an http client.
		if t >= 0 {
			options.Timeout = t
		}
	})
}

// WithBackoffStrategy controls the wait time between requests when retrying.
func WithBackoffStrategy(strategy BackoffFunc) Option {
	return optFunc(func(options *Options) {
		options.BackoffStrategy = strategy
	})
}

// WithRetryPolicy controls the retry policy of the http client.
func WithRetryPolicy(checkRetry CheckRetryFunc) Option {
	return optFunc(func(options *Options) {
		options.CheckRetry = checkRetry
	})
}

var (
	// defaultRetryMax is the maximum number of retries used by default.
	defaultRetryMax = 3

	// defaultHTTPClient is the http client used by default on requests.
	defaultHTTPClient = &http.Client{Timeout: defaultTimeout}

	// defaultTimeout is the timeout used by default on http client.
	// If a custom http client is provided and that http client has
	// a defined timeout, it will be overrided by defaultTimeout.
	// To set custom http client timeout, a custom timeout should
	// be provided also.
	defaultTimeout = 10 * time.Second

	// defaultBackoffStrategy is the retry strategy used by default by
	// the http client.
	defaultBackoffStrategy = ConstantBackoff(time.Second * 2)

	// defaultRetryPolicy is the function that tells on any given request if the
	// http client should retry it or not. By default, it retries on connection and 5xx errors only.
	defaultRetryPolicy = ServerErrorsRetryPolicy()
)

// DefaultOptions returns the default options.
func DefaultOptions() Options {
	return Options{
		RetryMax:        defaultRetryMax,
		HTTPClient:      defaultHTTPClient,
		Timeout:         defaultTimeout,
		BackoffStrategy: defaultBackoffStrategy,
		CheckRetry:      defaultRetryPolicy,
	}
}
