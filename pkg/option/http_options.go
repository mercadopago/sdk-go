package option

import (
	"context"
	"net/http"
	"time"
)

// CheckRetryFunc specifies a policy for handling retries. It is called
// following each request with the response and error values returned by
// the http.Client. If CheckRetryFunc returns false, the http client stops retrying
// and returns the response to the caller. If CheckRetryFunc returns an error,
// that error value is returned in lieu of the error from the request. The
// Client will close any response body when retrying, but if the retry is
// aborted it is up to the CheckResponse callback to properly close any
// response body before returning.
type CheckRetryFunc func(ctx context.Context, resp *http.Response, err error) (bool, error)

// BackoffFunc specifies a policy for how long to wait between retries. It is
// called after a failing request to determine the amount of time that should
// pass before trying again.
type BackoffFunc func(attempt int) time.Duration

type HTTPOptions struct {
	RetryMax int

	HTTPClient      *http.Client
	Timeout         time.Duration
	BackoffStrategy BackoffFunc
	CheckRetry      CheckRetryFunc
}

// HTTPOption signature for client configurable parameters.
type HTTPOption interface {
	ApplyHTTP(opts *HTTPOptions)
}

type optFunc func(opts *HTTPOptions)

func (f optFunc) ApplyHTTP(o *HTTPOptions) { f(o) }

// WithRetryMax tells the client the maximum number of retries to execute. Eg.: A
// value of 3, means to execute the original request, and up-to 3 retries (4
// requests in total). A value of 0 means no retries.
func WithRetryMax(max int) HTTPOption {
	return optFunc(func(options *HTTPOptions) {
		options.RetryMax = max
	})
}

// WithCustomClient allow do requests using received http client.
func WithCustomClient(c *http.Client) HTTPOption {
	return optFunc(func(options *HTTPOptions) {
		if c != nil {
			options.HTTPClient = c
		}
	})
}

// WithTimeout controls the timeout for each request. When retrying requests,
// each retried request will start counting from the beginning towards this
// timeout.
//
// A timeout of 0 disables request timeouts.
func WithTimeout(t time.Duration) HTTPOption {
	return optFunc(func(options *HTTPOptions) {
		// Negative durations do not make sense in the context of an http client.
		if t >= 0 {
			options.Timeout = t
		}
	})
}

// WithBackoffStrategy controls the wait time between requests when retrying.
func WithBackoffStrategy(strategy BackoffFunc) HTTPOption {
	return optFunc(func(options *HTTPOptions) {
		options.BackoffStrategy = strategy
	})
}

// WithRetryPolicy controls the retry policy of the http client.
func WithRetryPolicy(checkRetry CheckRetryFunc) HTTPOption {
	return optFunc(func(options *HTTPOptions) {
		options.CheckRetry = checkRetry
	})
}
