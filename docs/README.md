# Mercado Pago Go SDK documentation

Use the root README to install the SDK and send a first request. This index connects runnable examples with the operational behavior implemented by the Go SDK.

## Start with Orders

Orders is the recommended learning path because it demonstrates the complete payment lifecycle. Begin with [create](../examples/apis/order/create/main.go), inspect the result with [get](../examples/apis/order/get/main.go), and continue with [process](../examples/apis/order/process/main.go), [capture](../examples/apis/order/capture/main.go), [cancel](../examples/apis/order/cancel/main.go), or a [total refund](../examples/apis/order/refundtotal/main.go).

Add [request deadlines](configuration.md) before production traffic, then apply [retries and idempotency](retries-and-idempotency.md) and [typed error handling](error-handling.md). The [Orders-first example catalog](../examples/README.md#orders-first) also covers transactions, partial refunds, PSE, 3DS, automatic payments, and industry fields.

## Operational guides

| Question | Guide |
|---|---|
| How should clients and deadlines be configured? | [Configuration](configuration.md) |
| How should failures be classified? | [Error handling](error-handling.md) |
| When is a repeated write safe? | [Retries and idempotency](retries-and-idempotency.md) |
| How should signed notifications be processed? | [Webhook validation](webhooks.md) |
| What should be checked first during an incident? | [Troubleshooting](troubleshooting.md) |

Endpoint fields and response schemas belong to the [Mercado Pago API reference](https://www.mercadopago.com/developers/en/reference).
