# Handle API and transport errors

Use `errors.As` from the most specific recoverable error to the common `*mperror.ResponseError`. Typed errors preserve the API response and still unwrap to that base error.

```go
resource, err := client.Create(ctx, request)
if err != nil {
	var rateLimit *mperror.RateLimitError
	var responseError *mperror.ResponseError

	switch {
	case errors.As(err, &rateLimit):
		log.Printf("rate limited; retry-after=%d", rateLimit.RetryAfter)
	case errors.As(err, &responseError):
		log.Printf("status=%d message=%s", responseError.StatusCode, responseError.Message)
	case errors.Is(err, context.DeadlineExceeded):
		log.Print("operation deadline exceeded")
	default:
		log.Print("transport or unexpected SDK failure")
	}
}
```

Choose recovery by category:

| Error | Application decision |
|---|---|
| Authentication or forbidden | Stop retrying; correct credentials or permissions. |
| Validation or bad request | Correct the request; an identical retry will fail again. |
| Idempotency | Reuse the key only for the original operation and reconcile its result. |
| Rate limit | Respect `RetryAfter` when present and reduce concurrency. |
| Dependency, resource locked, or server | Retry only within a bounded policy. |
| Context or network failure | The outcome may be unknown; query the resource before repeating a write. |

Log the HTTP status, SDK error type, request ID when available, operation name, and your own correlation ID. Do not log access tokens, card data, full request bodies, or complete API responses.
