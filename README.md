<div align="center">
  <img src="assets/mercado-pago.svg" alt="Mercado Pago" width="120" />
  <h1>Mercado Pago SDK for Go</h1>
  <p><strong>The official Mercado Pago SDK for Go applications.</strong></p>
  <p>Build payment experiences with a typed, idiomatic client for the Mercado Pago API.</p>
  <p>
    <a href="https://github.com/mercadopago/sdk-go/actions/workflows/ci.yml"><img src="assets/badges/ci.svg" alt="CI workflow" /></a>
    <a href="#installation"><img src="assets/badges/package.svg" alt="Go package" /></a>
    <img src="assets/badges/runtime.svg" alt="Go 1.23 or later" />
    <a href="LICENSE"><img src="assets/badges/license.svg" alt="MIT license" /></a>
  </p>
  <p>
    <a href="#installation">Installation</a> ·
    <a href="#getting-started">Getting started</a> ·
    <a href="#documentation">Documentation</a> ·
    <a href="#contributing">Contributing</a>
  </p>
</div>

---

## 💡 Requirements

The SDK requires Go 1.23 or higher.

## 📲 Installation

First time using Mercado Pago? Create your [Mercado Pago account](https://www.mercadopago.com), if you don't have one already.

Install the Mercado Pago SDK for Go:

```sh
$ go install github.com/mercadopago/sdk-go
```

That's it! The Mercado Pago SDK for Go has been successfully installed.

## 🌟 Getting Started

Simple usage looks like:

```go
package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "online",
		TotalAmount:       "1000.00",
		ExternalReference: "ext_ref_1234",
		Currency:          "BRL",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "master",
						Token:        "{{CARD_TOKEN}}",
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		},
		Payer: &order.PayerRequest{
			Email: "{{PAYER_EMAIL}}",
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resource)
}
```

### SDK Configuration

Before making API requests, you need to initialize the SDK with your access token:

```go
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
```

### Making API Requests

To make requests to the Mercado Pago APIs, you can use the packages provided by the SDK. For example, to list payment methods, you can use the `paymentmethod` package:

```go
	client := paymentmethod.NewClient(cfg)
	resources, err := client.List(context.Background())
```

### Setting a custom idempotency key

For non-GET requests (Create, Cancel, Capture, etc.) the SDK automatically generates a UUID as the `X-Idempotency-Key` header. To supply your own stable key — for example, when retrying a failed request — attach it to the context using `requestoptions.WithIdempotencyKey` before calling the client method:

```go
import "github.com/mercadopago/sdk-go/pkg/requestoptions"

ctx := requestoptions.WithIdempotencyKey(context.Background(), "my-stable-idempotency-key")
resource, err := client.Create(ctx, request)
```

The key travels with the context and is applied automatically; no change to the method signature is required.

### Exception throwing handling

Every package methods returns two variables: response (type of the package) and error (type of the std lib), which will contain any error thrown. It is important to handle these errors in the best possible way.

```go
	resources, err := client.List(context.Background())
	if err != nil {
		// appropriate treatment
	}
```

For more details on available methods and request parameters, see the [SDK source packages](https://github.com/mercadopago/sdk-go/tree/main/pkg) and the documentation index below.

## 📚 Documentation

- [SDK guides](docs/README.md): configuration, errors, webhooks, and production guidance.
- [Examples](examples/README.md): integration flows supported by this SDK.
- [Go source packages](https://github.com/mercadopago/sdk-go/tree/main/pkg): public packages and methods.
- [API reference](https://www.mercadopago.com/developers/en/reference): request fields and response schemas.

## 🤝 Contributing

All contributions are welcome, ranging from people wanting to triage issues, others wanting to write documentation, to people wanting to contribute with code.

Please read and follow our [contribution guidelines](CONTRIBUTING.md). Contributions not following these guidelines will be disregarded. The guidelines are in place to make all of our lives easier and make contribution a consistent process for everyone.

## ❤️ Support

If you require technical support, please contact our support team at our developers site: [English](https://www.mercadopago.com/developers/en/support/center/contact) | [Portuguese](https://www.mercadopago.com/developers/pt/support/center/contact) | [Spanish](https://www.mercadopago.com/developers/es/support/center/contact)
