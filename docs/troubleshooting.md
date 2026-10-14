# Troubleshooting

## A request times out repeatedly

Check both `config.WithTimeout` and the caller's context deadline. The per-attempt timeout should be shorter than the total operation deadline so retries have time to run. If a write timed out, reconcile the resource before repeating it with the same idempotency key.

## A custom HTTP client ignores SDK settings

Do not combine `WithHTTPClient` with `WithTimeout` or `WithMaxRetries`: current configuration construction gives those default-requester options precedence. Configure timeout, retries, proxy, TLS, and connection pooling inside the custom requester.

## The API keeps returning `429`

Read `mperror.RateLimitError.RetryAfter`, reduce concurrency, and avoid immediate application-level retry loops around the SDK's built-in retries.

## A webhook signature is rejected

Capture the raw values of `x-request-id` and `data.id` without logging secrets. Check clock synchronization and the configured tolerance. Do not pass unsigned QR Code notifications to the validator.

## Escalating an issue

Record the SDK version, operation, HTTP status, request ID, sanitized error type/message, and a minimal reproducible example. Open a repository issue without credentials, card data, webhook secrets, or exploit details.
