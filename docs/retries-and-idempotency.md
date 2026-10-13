# Retries and idempotency

Retries solve transient delivery failures; idempotency prevents a retried write from creating the same business operation twice. Use both deliberately.

The default requester retries network failures and HTTP `429`, `500`, `502`, `503`, and `504`, up to three attempts with a two-second delay. A context deadline can end the sequence earlier. Validation, authentication, and other permanent failures are not retryable.

Attach a stable key to every retryable write:

```go
ctx := requestoptions.WithIdempotencyKey(
	context.Background(),
	"checkout-order-2026-0001",
)

payment, err := paymentClient.Create(ctx, request)
```

If no key is supplied, the SDK generates one. Supplying your own lets the application persist it and safely continue the same operation after a process restart.

Use this lifecycle:

1. Generate the key when the business operation is created.
2. Persist it with your internal order or command.
3. Reuse it only while attempting that exact operation with the same payload.
4. After a timeout or network error, query by your external reference or known resource ID before sending the write again.
5. Generate a new key for a genuinely new operation.

Do not create a fresh key inside a retry loop. That turns each attempt into a new operation and defeats duplicate protection.
