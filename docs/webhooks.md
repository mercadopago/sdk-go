# Validate and process webhooks

Validate signed notifications before trusting identifiers or scheduling business work:

```go
err := webhook.ValidateSignature(
	r.Header.Get("x-signature"),
	r.Header.Get("x-request-id"),
	r.URL.Query().Get("data.id"),
	os.Getenv("MP_WEBHOOK_SECRET"),
	webhook.WithTolerance(5*time.Minute),
)
if err != nil {
	var signatureError *webhook.SignatureError
	if errors.As(err, &signatureError) {
		http.Error(w, "invalid signature", http.StatusUnauthorized)
		return
	}
	http.Error(w, "invalid notification", http.StatusBadRequest)
	return
}
```

Pass `x-signature`, `x-request-id`, and `data.id` exactly as received. URL decoding, trimming, or changing case before validation can invalidate the manifest. QR Code notifications are not signed and must not use this validator.

A production handler should:

1. Read the required headers and query parameter.
2. Validate the signature and timestamp tolerance.
3. Deduplicate using a notification or resource identifier.
4. Enqueue the identifier and return a successful response promptly.
5. Fetch the current resource from Mercado Pago inside the worker.
6. Apply state transitions idempotently; notifications may be repeated or arrive out of order.

Treat the notification as a signal, not as the authoritative resource state. Store the outcome and reason of validation, but never log the webhook secret or complete signature.
