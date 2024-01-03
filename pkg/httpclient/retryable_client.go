package httpclient

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"time"
)

type retryAttemptContextKey struct{}

// CheckRetryFunc specifies a policy for handling retries. It is called
// following each request with the response and error values returned by
// the http.Client. If CheckRetryFunc returns false, the Client stops retrying
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

// RetryableClient is a compatible http.Client that allows the caller to setup
// a retry strategy for retrying failed requests transparently.
type RetryableClient struct {
	// Compose a *http.Client, we'll override the Do method.
	*http.Client

	// CheckRetry specifies the policy for handling retries. It is called after
	// each request. If CheckRetryFunc is nil then HTTPRetryPolicy will be used.
	CheckRetry CheckRetryFunc

	// RetryMax is the maximum number of retries to do before returning an error.
	RetryMax int

	// BackoffStrategy tells the client how much time it must wait between retries.
	BackoffStrategy BackoffFunc
}

// Do sends an HTTP request and returns an HTTP response, following policy
// (such as redirects, cookies, auth) as configured on the client.
func (c *RetryableClient) Do(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; ; i++ {
		req, err = requestFromInternal(req, i)
		if err != nil {
			return nil, err
		}

		// Attempt the request using the underlying httpClient.
		resp, err = c.Client.Do(req)

		// Check if we should continue with retries. We always check after a request
		// to allow the user to define what a successful request is. If this call
		// return (false, nil) then we can assert that the request was successful
		// and therefore, we can return the given response to the user.
		shouldRetry, retryErr := c.checkRetry(req.Context(), resp, err)

		// Now decide if we should continue based on checkRetries answer.
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
			c.drainBody(resp.Body)
		}

		// Call Backoff to see how much time we must wait until next retry.
		backoffWait := c.backoffDuration(i, resp)

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
	// If this is a retry attempt then set a value in context indicating so.
	// This is handy for clients willing to insert custom headers.
	ctx := req.Context()
	if retryAttempt > 0 {
		ctx = withRetries(ctx, retryAttempt)
	}

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

// withRetries returns a new context decorated with a retry count.
func withRetries(ctx context.Context, retryAttempt int) context.Context {
	return context.WithValue(ctx, retryAttemptContextKey{}, retryAttempt)
}

func (c *RetryableClient) checkRetry(ctx context.Context, res *http.Response, err error) (bool, error) {
	if c.CheckRetry != nil {
		return c.CheckRetry(ctx, res, err)
	}
	return ServerErrorsRetryPolicy()(ctx, res, err)
}

// ServerErrorsRetryPolicy provides a sane default implementation of a
// CheckRetryFunc, it will retry on server (5xx) errors.
func ServerErrorsRetryPolicy() CheckRetryFunc {
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
func (c *RetryableClient) drainBody(body io.ReadCloser) {
	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	const respReadLimit = int64(4096)

	defer body.Close()
	_, _ = io.Copy(io.Discard, io.LimitReader(body, respReadLimit))
}

func (c *RetryableClient) backoffDuration(attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := retryAfterDuration(s[0]); err == nil {
					return sleep
				}
			}
		}
	}

	if c.BackoffStrategy != nil {
		return c.BackoffStrategy(attemptNum)
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

// ConstantBackoff provides a callback for Client.Backoff which will perform
// linear backoff based on the provided minimum duration.
func ConstantBackoff(wait time.Duration) BackoffFunc {
	return func(_ int) time.Duration {
		return wait
	}
}
