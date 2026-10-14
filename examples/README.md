# Go examples

These programs demonstrate integration flows implemented by the Go SDK. Start with Orders, then choose another API only when the payment product requires it.

## Orders first

Orders is the most complete learning path in this repository: it covers creation, retrieval, processing, capture, cancellation, refunds, and transaction management.

| Stage | Goal | Example |
|---|---|---|
| 1. Create | Submit a basic order | [`create/`](apis/order/create/) |
| 2. Inspect | Retrieve the server-side result | [`get/`](apis/order/get/) |
| 3. Advance | Process or capture an order | [`process/`](apis/order/process/) · [`capture/`](apis/order/capture/) |
| 4. Reverse | Cancel, partially refund, or fully refund | [`cancel/`](apis/order/cancel/) · [`refundpartial/`](apis/order/refundpartial/) · [`refundtotal/`](apis/order/refundtotal/) |
| 5. Manage transactions | Create, update, or delete a transaction | [`createtransaction/`](apis/order/createtransaction/) · [`updatetransaction/`](apis/order/updatetransaction/) · [`delete/`](apis/order/delete/) |
| 6. Add a payment flow | Use PSE, 3DS, automatic payments, or industry fields | [`createpse/`](apis/order/createpse/) · [`create3ds/`](apis/order/create3ds/) · [`createwithautomaticpayments/`](apis/order/createwithautomaticpayments/) · [`createwithindustryfields/`](apis/order/createwithindustryfields/) |

Run an example from the repository root with `go run ./examples/apis/order/create`. Before production use, add a context deadline, persist the idempotency key, and handle typed errors. See [configuration](../docs/configuration.md), [retries and idempotency](../docs/retries-and-idempotency.md), and [error handling](../docs/error-handling.md).

## Other supported examples

| Goal | API | Examples |
|---|---|---|
| Create a hosted checkout | Preferences API | [`apis/preference/`](apis/preference/) |
| Connect a seller account | OAuth API | [`apis/oauth/`](apis/oauth/) |
| Create or manage subscriptions | Preapproval APIs | [`apis/preapproval/`](apis/preapproval/) · [`apis/preapprovalplan/`](apis/preapprovalplan/) |
| Create and manage payments | Payments API | [`apis/payment/`](apis/payment/) |
| Create a Point payment intent | Point API | [`apis/point/`](apis/point/) |
| Customize transport and timeout | SDK configuration | [`custom_http_client/`](custom_http_client/) · [`custom_timeout/`](custom_timeout/) |
| Parse API failures | SDK errors | [`mp_error_parse/`](mp_error_parse/) |

Use test credentials and test data, replace every placeholder, and confirm fields in the [API reference](https://www.mercadopago.com/developers/en/reference). Never commit credentials or use production card data in examples.
