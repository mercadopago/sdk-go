package rest

import (
	"math"
	"net/http"
	"time"
)

const (
	defaultMaxRetries = 5
	defaultMaxBackoff = time.Minute
	defaultRetryDelay = time.Second
)

// RetryClient is the interface that defines retry signature.
type RetryClient interface {
	Retry(req *http.Request, httpClient *http.Client, opts ...Option) (*http.Response, error)
}

// retryClient is the default implementation of RetryClient.
type retryClient struct{}

func (*retryClient) Retry(req *http.Request, httpClient *http.Client, opts ...Option) (*http.Response, error) {
	var (
		maxRetries = defaultMaxRetries
		maxBackoff = defaultMaxBackoff
		retryDelay = defaultRetryDelay
		res        *http.Response
		err        error
	)

	options := &options{}
	for _, opt := range opts {
		opt.apply(options)
	}
	if options.maxRetries > 0 {
		maxRetries = options.maxRetries
	}
	if options.maxBackoff > 0 {
		maxBackoff = options.maxBackoff
	}
	if options.retryDelay > 0 {
		retryDelay = options.retryDelay
	}

	for i := 0; i < maxRetries; i++ {
		time.Sleep(retryDelay)

		res, err = httpClient.Do(req)
		if shouldStop(res, err) {
			break
		}

		newDelay := float64(retryDelay * 2)
		retryDelay = time.Duration(math.Min(newDelay, float64(maxBackoff)))
	}

	return res, err
}

func shouldStop(res *http.Response, err error) bool {
	return err == nil && res.StatusCode < http.StatusInternalServerError
}
