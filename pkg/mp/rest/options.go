package rest

import (
	"net/http"
	"time"
)

type options struct {
	maxRetries int

	maxBackoff    time.Duration
	retryDelay    time.Duration
	timeout       time.Duration
	customHeaders http.Header
}

type Option interface {
	apply(*options)
}

type maxRetriesOption int

func (m maxRetriesOption) apply(opts *options) {
	opts.maxRetries = int(m)
}

func WithMaxRetries(m int) Option {
	return maxRetriesOption(m)
}

type maxBackoffOption time.Duration

func (m maxBackoffOption) apply(opts *options) {
	opts.maxBackoff = time.Duration(m)
}

func WithMaxBackoff(t time.Duration) Option {
	return maxBackoffOption(t)
}

type retryDelayOption time.Duration

func (r retryDelayOption) apply(opts *options) {
	opts.retryDelay = time.Duration(r)
}

func WithRetryDelay(t time.Duration) Option {
	return retryDelayOption(t)
}

type timeoutOption time.Duration

func (t timeoutOption) apply(opts *options) {
	opts.timeout = time.Duration(t)
}

func WithTimeout(t time.Duration) Option {
	return timeoutOption(t)
}

type customHeadersOption http.Header

func (c customHeadersOption) apply(opts *options) {
	opts.customHeaders = http.Header(c)
}

func WithCustomHeaders(h http.Header) Option {
	return customHeadersOption(h)
}
