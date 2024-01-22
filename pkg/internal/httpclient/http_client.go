package httpclient

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/mercadopago/sdk-go/pkg/option"
)

// do sends an HTTP request and returns an HTTP response, following policy
// (such as redirects, cookies, auth) as configured on the http client.
func do(ctx context.Context, req *http.Request, c option.HTTPOptions) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; ; i++ {
		req, err = requestFromInternal(req, i)
		if err != nil {
			return nil, err
		}

		// Attempt the request using the underlying http client.
		resp, err = c.HTTPClient.Do(req)

		// Check if we should continue with retries. We always check after a request
		// to allow the user to define what a successful request is. If this call
		// return (false, nil) then we can assert that the request was successful
		// and therefore, we can return the given response to the user.
		shouldRetry, retryErr := shouldRetry(ctx, c.CheckRetry, resp, err)

		// Now decide if we should continue based on shouldRetry answer.
		if !shouldRetry {
			if retryErr != nil {
				err = retryErr
			}
			return resp, err
		}

		// If we have no retries left then we return the last response and error
		// from the last request executed by the client.
		remainingRetries := c.RetryMax - i
		if remainingRetries <= 0 {
			return resp, err
		}

		// We're going to retry, consume any response so TCP connection can be reused.
		if err == nil && resp != nil {
			drainBody(resp.Body)
		}

		// Call Backoff to see how much time we must wait until next retry.
		backoffWait := backoffDuration(i, c.BackoffStrategy, resp)

		// If the request context has a deadline, check whether that deadline
		// happens before the wait period of the backoff strategy. In case
		// it does we return the last error without waiting.
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
func requestFromInternal(req *http.Request, retryAttempt int) (*http.Request, error) {
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

func shouldRetry(ctx context.Context, checkRetry option.CheckRetryFunc, res *http.Response, err error) (bool, error) {
	if checkRetry != nil {
		return checkRetry(ctx, res, err)
	}
	return serverErrorsRetryPolicy()(ctx, res, err)
}

// serverErrorsRetryPolicy provides a sane default implementation of a
// CheckRetryFunc, it will retry on server (5xx) errors.
func serverErrorsRetryPolicy() option.CheckRetryFunc {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
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
}

// Try to read the response body so we can reuse this connection.
func drainBody(body io.ReadCloser) {
	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	const respReadLimit = int64(4096)

	defer body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(body, respReadLimit))
}

func backoffDuration(attemptNum int, backoffStrategy option.BackoffFunc, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := retryAfterDuration(s[0]); err == nil {
					return sleep
				}
			}
		}
	}

	if backoffStrategy != nil {
		return backoffStrategy(attemptNum)
	}

	return 0
}

// retryAfterDuration returns the duration for the Retry-After header.
func retryAfterDuration(t string) (time.Duration, error) {
	when, err := time.Parse(http.TimeFormat, t)
	if err == nil {
		// when is always in UTC, so make the math from UTC+0.
		t := time.Now().UTC()
		return when.Sub(t), nil
	}

	// The duration can be in seconds.
	d, err := strconv.Atoi(t)
	if err != nil {
		return 0, err
	}

	return time.Duration(d) * time.Second, nil
}

// ConstantBackoff provides a callback for backoffStrategy which will perform
// linear backoff based on the provided minimum duration.
func ConstantBackoff(wait time.Duration) option.BackoffFunc {
	return func(_ int) time.Duration {
		return wait
	}
}
