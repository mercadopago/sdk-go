package httpclient

import (
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/option"
)

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
func DefaultOptions() option.HTTPOptions {
	return option.HTTPOptions{
		RetryMax:        defaultRetryMax,
		HTTPClient:      defaultHTTPClient,
		Timeout:         defaultTimeout,
		BackoffStrategy: defaultBackoffStrategy,
		CheckRetry:      defaultRetryPolicy,
	}
}
