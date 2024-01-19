package header

import (
	"context"
	"net/http"
)

type customHeadersCtxKey struct{}

func Context(ctx context.Context, customHeaders http.Header) context.Context {
	return context.WithValue(ctx, customHeadersCtxKey{}, customHeaders)
}

func Headers(ctx context.Context) http.Header {
	value, _ := ctx.Value(customHeadersCtxKey{}).(http.Header)
	return value
}
