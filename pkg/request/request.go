package request

import (
	"context"
	"net/http"
)

type customClientCtxKey struct{}

func WithCustomClient(ctx context.Context, customClient *http.Client) context.Context {
	return context.WithValue(ctx, customClientCtxKey{}, customClient)
}

func CustomClient(ctx context.Context) *http.Client {
	value, _ := ctx.Value(customClientCtxKey{}).(*http.Client)
	return value
}

type customHeadersCtxKey struct{}

func WithCustomHeaders(ctx context.Context, customHeaders http.Header) context.Context {
	return context.WithValue(ctx, customHeadersCtxKey{}, customHeaders)
}

func CustomHeaders(ctx context.Context) http.Header {
	value, _ := ctx.Value(customHeadersCtxKey{}).(http.Header)
	return value
}
