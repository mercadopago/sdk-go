# Configure the Go SDK

Create one configuration at application startup and reuse it when building resource clients:

```go
cfg, err := config.New(
	os.Getenv("MP_ACCESS_TOKEN"),
	config.WithTimeout(10*time.Second),
	config.WithMaxRetries(3),
)
if err != nil {
	log.Fatal(err)
}
```

`WithTimeout` limits each HTTP attempt. Also give every operation a context deadline so the complete call—including retries—cannot outlive the request or job that started it:

```go
ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
defer cancel()

payment, err := paymentClient.Get(ctx, paymentID)
```

The default requester makes at most three retry attempts with a constant two-second delay. It retries network failures and HTTP `429`, `500`, `502`, `503`, and `504`; context cancellation and deadlines stop the retry loop.

`WithHTTPClient` replaces the default requester when used on its own. In the current implementation, combining it with `WithTimeout` or `WithMaxRetries` rebuilds the SDK default requester, so configure timeout, retry, proxy, TLS, and connection pooling inside the custom requester instead of mixing those options. `WithCorporationID`, `WithIntegratorID`, and `WithPlatformID` add partner metadata only when those identifiers were assigned to your integration.

Keep credentials in environment variables or a secret manager. Never create a configuration from an access token supplied by an end user, and never log the configuration object.
