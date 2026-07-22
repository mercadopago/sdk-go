// Package requestoptions provides per-request configuration for SDK clients
// via [context.Context]. Use [WithIdempotencyKey] to attach a caller-controlled
// idempotency key to a context before passing it to any mutating client method.
package requestoptions

import "context"

type contextKey struct{}

// WithIdempotencyKey returns a new context that carries the given idempotency
// key. Pass the resulting context to any mutating client method (Create, Cancel,
// Capture, etc.) to override the auto-generated UUID that the SDK would
// otherwise set in the X-Idempotency-Key header:
//
//	ctx = requestoptions.WithIdempotencyKey(ctx, "my-stable-key")
//	resp, err := client.Create(ctx, req)
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, contextKey{}, key)
}

// IdempotencyKeyFrom returns the idempotency key stored in ctx by
// [WithIdempotencyKey], and whether one was set.
func IdempotencyKeyFrom(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	key, ok := ctx.Value(contextKey{}).(string)
	return key, ok && key != ""
}
