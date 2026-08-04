// Package defaultrequester provides the SDK's built-in [requester.Requester]
// implementation. It wraps the standard [net/http.Client] with automatic retry
// logic for transient server errors (HTTP 5xx, excluding 501 Not Implemented)
// and network-level failures. Retries use a constant back-off strategy and
// respect the request context's deadline and cancellation signals.
//
// This package is internal; SDK consumers interact with it indirectly through
// [config.New] or by supplying their own [requester.Requester] via
// [config.WithHTTPClient].
package defaultrequester

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/mercadopago/sdk-go/pkg/requester"
)

var (
	// defaultRetryMax is the maximum number of retry attempts before the
	// requester gives up and returns the last error.
	defaultRetryMax = 3

	// defaultHTTPClient is the shared [http.Client] used for all requests made
	// through the default requester. It enforces [defaultTimeout] per request.
	defaultHTTPClient = &http.Client{Timeout: defaultTimeout}

	// defaultTimeout is the per-request timeout applied to the shared
	// [http.Client].
	defaultTimeout = 10 * time.Second

	// defaultBackoffStrategy defines how long to wait between consecutive retry
	// attempts. The default is a constant 2-second delay regardless of the
	// attempt number.
	defaultBackoffStrategy = constantBackoff(time.Second * 2)
)

// defaultRequester is a [requester.Requester] implementation that delegates HTTP
// execution with automatic retry support. Fields override the package defaults.
type defaultRequester struct {
	httpClient *http.Client
	maxRetries int
}

func (d *defaultRequester) client() *http.Client {
	if d.httpClient != nil {
		return d.httpClient
	}
	return defaultHTTPClient
}

func (d *defaultRequester) retryMax() int {
	if d.maxRetries > 0 {
		return d.maxRetries
	}
	return defaultRetryMax
}

// backoffFunc defines a strategy for computing the wait duration between retry
// attempts. It receives the zero-based attempt number and returns how long to
// sleep before the next try.
type backoffFunc func(attempt int) time.Duration

// New returns a new [requester.Requester] backed by the SDK's default HTTP
// client with automatic retries and constant back-off.
func New() requester.Requester {
	return &defaultRequester{}
}

// NewWithOptions returns a [requester.Requester] whose timeout and retry count
// are overridden by the supplied values. Zero values use the package defaults.
func NewWithOptions(timeout time.Duration, maxRetries int) requester.Requester {
	d := &defaultRequester{}
	if timeout > 0 {
		d.httpClient = &http.Client{Timeout: timeout}
	}
	if maxRetries > 0 {
		d.maxRetries = maxRetries
	}
	return d
}

// Do executes the given HTTP request with automatic retries on transient errors
// (429, 5xx excluding 501, and network failures). It rewinds the request body
// between attempts, respects context deadlines/cancellations, and drains
// response bodies of failed attempts so TCP connections can be reused.
func (d *defaultRequester) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; ; i++ {
		req, err = requestFromInternal(req)
		if err != nil {
			return nil, err
		}

		resp, err = d.client().Do(req)

		shouldRetryResult, retryErr := shouldRetry(req.Context(), resp, err)
		if !shouldRetryResult {
			if retryErr != nil {
				err = retryErr
			}
			return resp, err
		}

		remainingRetries := d.retryMax() - i
		if remainingRetries <= 0 {
			return resp, err
		}

		if err == nil && resp != nil {
			drainBody(resp.Body)
		}

		backoffWait := backoffDuration(i)

		if deadline, ok := req.Context().Deadline(); ok {
			ctxDeadline := time.Until(deadline)
			if ctxDeadline <= backoffWait {
				return resp, err
			}
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(backoffWait):
		}
	}
}

// requestFromInternal builds an *http.Request from our internal request.
func requestFromInternal(req *http.Request) (*http.Request, error) {
	ctx := req.Context()

	// Use the context from the internal request. When cloning requests
	// we want to have the same context in all of them. The request
	// might pass through a number of hooks which are allowed
	// to change its context.
	r2 := req.WithContext(ctx)

	// Always rewind the request body when non-nil.
	if req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		r2.Body = body
	}

	return r2, nil
}

// defaultRetryOn contains the HTTP status codes that trigger automatic retries.
// Includes 429 (Too Many Requests) and 5xx server errors (excluding 501).
var defaultRetryOn = map[int]bool{
	429: true,
	500: true,
	502: true,
	503: true,
	504: true,
}

// shouldRetry decides whether a failed request should be retried.
// Retries on: network errors, 429 Too Many Requests, and 5xx server errors (excluding 501).
func shouldRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		return true, err
	}

	// Retry on codes in defaultRetryOn (429 + 5xx except 501)
	if defaultRetryOn[resp.StatusCode] {
		return true, nil
	}
	// Also catch unexpected codes (0, 999, etc.)
	if resp.StatusCode == 0 {
		return true, nil
	}

	return false, nil
}

// Try to read the response body, so we can reuse this connection.
func drainBody(body io.ReadCloser) {
	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	const respReadLimit = int64(4096)

	defer body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(body, respReadLimit))
}

func backoffDuration(attemptNum int) time.Duration {
	return defaultBackoffStrategy(attemptNum)
}

// constantBackoff provides a callback for backoffStrategy which will perform
// linear backoff based on the provided minimum duration.
func constantBackoff(wait time.Duration) backoffFunc {
	return func(_ int) time.Duration {
		return wait
	}
}
