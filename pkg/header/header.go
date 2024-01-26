package header

import (
	"context"
	"net/http"
)

type customHeadersCtxKey string

const (
	customHeadersKey = customHeadersCtxKey("customHeaders")
)

func Context(ctx context.Context, customHeaders http.Header) context.Context {
	return context.WithValue(ctx, customHeadersKey, customHeaders)
}

func Headers(ctx context.Context) http.Header {
	value, _ := ctx.Value(customHeadersKey).(http.Header)
	return value
}
