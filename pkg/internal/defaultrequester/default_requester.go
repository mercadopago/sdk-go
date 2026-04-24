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

// defaultRequester is a stateless [requester.Requester] implementation that
// delegates HTTP execution to [defaultHTTPClient] with automatic retry support.
type defaultRequester struct{}

// backoffFunc defines a strategy for computing the wait duration between retry
// attempts. It receives the zero-based attempt number and returns how long to
// sleep before the next try.
type backoffFunc func(attempt int) time.Duration

// New returns a new [requester.Requester] backed by the SDK's default HTTP
// client with automatic retries and constant back-off.
func New() requester.Requester {
	return &defaultRequester{}
}

// Do executes the given HTTP request with automatic retries on transient server
// errors (HTTP 5xx, excluding 501). It rewinds the request body between attempts,
// respects the request context's deadline and cancellation, and drains response
// bodies of failed attempts so that the underlying TCP connection can be reused.
func (d *defaultRequester) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; ; i++ {
		req, err = requestFromInternal(req)
		if err != nil {
			return nil, err
		}

		// Attempt the request using the default http client.
		resp, err = defaultHTTPClient.Do(req)

		// Check if we should continue with retries. We always check after a request
		// to allow the user to define what a successful request is. If this call
		// return (false, nil) then we can assert that the request was successful
		// and therefore, we can return the given response to the user.
		shouldRetry, retryErr := shouldRetry(req.Context(), resp, err)

		// Now decide if we should continue based on shouldRetry answer.
		if !shouldRetry {
			if retryErr != nil {
				err = retryErr
			}
			return resp, err
		}

		// If we have no retries left then we return the last response and error
		// from the last request executed by the client.
		remainingRetries := defaultRetryMax - i
		if remainingRetries <= 0 {
			return resp, err
		}

		// We're going to retry, consume any response so TCP connection can be reused.
		if err == nil && resp != nil {
			drainBody(resp.Body)
		}

		// Call backoff to see how much time we must wait until next retry.
		backoffWait := backoffDuration(i)

		// If the request context has a deadline, check whether that deadline
		// happens before the wait period of the backoff strategy. In case
		// it do we return the last error without waiting.
		if deadline, ok := req.Context().Deadline(); ok {
			ctxDeadline := time.Until(deadline)
			if ctxDeadline <= backoffWait {
				return resp, err
			}
		}

		// Wait for either the backoff period or the cancellation of the request context.
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

// shouldRetry provides a sane default implementation of a
// retry policy, it will retry on server (5xx) errors.
func shouldRetry(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err != nil {
		return true, err
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
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
