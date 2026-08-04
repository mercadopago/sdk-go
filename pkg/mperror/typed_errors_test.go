package mperror_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/mperror"
)

// ── Exception hierarchy ───────────────────────────────────────────────────────

func TestAllSubtypesEmbedResponseError(t *testing.T) {
	t.Parallel()

	// Use BuildError to create subtypes so we don't depend on internal struct layout.
	statuses := []int{400, 401, 402, 403, 404, 409, 422, 423, 424, 429, 500}
	subtypes := make([]error, len(statuses))
	for i, s := range statuses {
		base := &mperror.ResponseError{StatusCode: s, Headers: http.Header{}}
		subtypes[i] = mperror.BuildError(base, 0)
	}

	for _, err := range subtypes {
		var respErr *mperror.ResponseError
		if !errors.As(err, &respErr) {
			t.Errorf("errors.As(*ResponseError) failed for %T", err)
		}
	}
}

func TestRateLimitErrorStoresRetryAfter(t *testing.T) {
	t.Parallel()
	base := &mperror.ResponseError{StatusCode: 429, Headers: http.Header{}}
	err := mperror.BuildError(base, 45).(*mperror.RateLimitError)
	if err.RetryAfter != 45 {
		t.Errorf("expected RetryAfter=45, got %d", err.RetryAfter)
	}
}

// ── BuildError factory ─────────────────────────────────────────────────────

func TestBuildErrorFactory(t *testing.T) {
	t.Parallel()

	cases := []struct {
		status int
		target interface{ Error() string }
	}{
		{400, &mperror.BadRequestError{}},
		{401, &mperror.AuthenticationError{}},
		{402, &mperror.PaymentError{}},
		{403, &mperror.ForbiddenError{}},
		{404, &mperror.NotFoundError{}},
		{409, &mperror.IdempotencyError{}},
		{422, &mperror.ValidationError{}},
		{423, &mperror.ResourceLockedError{}},
		{424, &mperror.DependencyError{}},
		{429, &mperror.RateLimitError{}},
		{500, &mperror.ServerError{}},
		{503, &mperror.ServerError{}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(http.StatusText(tc.status), func(t *testing.T) {
			t.Parallel()
			base := &mperror.ResponseError{StatusCode: tc.status, Headers: http.Header{}}
			got := mperror.BuildError(base, 0)
			if !errors.As(got, &tc.target) {
				t.Errorf("status %d: expected %T, got %T", tc.status, tc.target, got)
			}
		})
	}
}

func TestBuildErrorUnknownClientStatusReturnsBase(t *testing.T) {
	t.Parallel()
	base := &mperror.ResponseError{StatusCode: 418, Headers: http.Header{}}
	got := mperror.BuildError(base, 0)
	var respErr *mperror.ResponseError
	if !errors.As(got, &respErr) {
		t.Errorf("expected *ResponseError for 418, got %T", got)
	}
}

func TestBuildError429WithRetryAfter(t *testing.T) {
	t.Parallel()
	base := &mperror.ResponseError{StatusCode: 429, Headers: http.Header{}}
	got := mperror.BuildError(base, 30)
	var rl *mperror.RateLimitError
	if !errors.As(got, &rl) {
		t.Fatalf("expected *RateLimitError, got %T", got)
	}
	if rl.RetryAfter != 30 {
		t.Errorf("expected RetryAfter=30, got %d", rl.RetryAfter)
	}
}

func TestBuildErrorNilBaseReturnsNil(t *testing.T) {
	t.Parallel()
	if mperror.BuildError(nil, 0) != nil {
		t.Error("expected nil for nil base")
	}
}

// ── 429 in shouldRetry (integration via defaultrequester package is in its own test) ─
