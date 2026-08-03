package mperror

// Typed error subtypes for the MercadoPago Go SDK.
//
// Each subtype embeds typedBase (which wraps *ResponseError + Unwrap) so that
// existing errors.As assertions on *ResponseError continue to work:
//
//	var respErr *mperror.ResponseError
//	if errors.As(err, &respErr) { ... }  // still works for all subtypes
//
// Specific type assertions also work:
//
//	var notFound *mperror.NotFoundError
//	if errors.As(err, &notFound) { ... }

// typedBase carries the underlying *ResponseError and exposes Unwrap() so that
// errors.As can traverse the chain to *ResponseError.
type typedBase struct{ *ResponseError }

// Unwrap returns the underlying *ResponseError for errors.As traversal.
func (b typedBase) Unwrap() error { return b.ResponseError }

// wrap is a convenience constructor used by the factory.
func wrap(r *ResponseError) typedBase { return typedBase{r} }

// BadRequestError represents HTTP 400 Bad Request (validation or syntax error).
type BadRequestError struct{ typedBase }

// AuthenticationError represents HTTP 401 Unauthorized (missing or invalid credentials).
type AuthenticationError struct{ typedBase }

// PaymentError represents HTTP 402 Payment Required (transaction processing error, AP/Orders).
type PaymentError struct{ typedBase }

// ForbiddenError represents HTTP 403 Forbidden.
type ForbiddenError struct{ typedBase }

// NotFoundError represents HTTP 404 Not Found.
type NotFoundError struct{ typedBase }

// IdempotencyError represents HTTP 409 Conflict (idempotency-key or state-machine conflict).
type IdempotencyError struct{ typedBase }

// ValidationError represents HTTP 422 Unprocessable Entity (business-rule violation).
type ValidationError struct{ typedBase }

// ResourceLockedError represents HTTP 423 Locked (idempotency key temporarily locked; retryable).
type ResourceLockedError struct{ typedBase }

// DependencyError represents HTTP 424 Failed Dependency (internal dependency failure; retryable).
type DependencyError struct{ typedBase }

// RateLimitError represents HTTP 429 Too Many Requests.
// RetryAfter holds the seconds to wait from the Retry-After header, or 0 if absent.
type RateLimitError struct {
	typedBase
	RetryAfter int
}

// ServerError represents HTTP 5xx Server Error.
type ServerError struct{ typedBase }

// statusMap maps exact HTTP status codes to their typed error constructor.
var statusMap = map[int]func(*ResponseError) error{
	400: func(r *ResponseError) error { return &BadRequestError{wrap(r)} },
	401: func(r *ResponseError) error { return &AuthenticationError{wrap(r)} },
	402: func(r *ResponseError) error { return &PaymentError{wrap(r)} },
	403: func(r *ResponseError) error { return &ForbiddenError{wrap(r)} },
	404: func(r *ResponseError) error { return &NotFoundError{wrap(r)} },
	409: func(r *ResponseError) error { return &IdempotencyError{wrap(r)} },
	422: func(r *ResponseError) error { return &ValidationError{wrap(r)} },
	423: func(r *ResponseError) error { return &ResourceLockedError{wrap(r)} },
	424: func(r *ResponseError) error { return &DependencyError{wrap(r)} },
}

// BuildError maps an HTTP status code to the most specific error subtype.
// The underlying *ResponseError is reachable via errors.As for existing catch patterns.
//
// retryAfter is only meaningful for status 429; pass 0 if Retry-After header was absent.
func BuildError(base *ResponseError, retryAfter int) error {
	if base == nil {
		return nil
	}
	if base.StatusCode == 429 {
		return &RateLimitError{typedBase: wrap(base), RetryAfter: retryAfter}
	}
	if ctor, ok := statusMap[base.StatusCode]; ok {
		return ctor(base)
	}
	if base.StatusCode >= 500 {
		return &ServerError{wrap(base)}
	}
	return base
}
