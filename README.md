# Mercado Pago SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mercadopago/dx-go.svg)](https://pkg.go.dev/github.com/mercadopago/dx-go)
[![License](https://img.shields.io/apm/l/vim-mode)](https://github.com/mercadopago/sdk-go)

A comprehensive Go client library for integrating with the Mercado Pago API.

## 💡 Requirements

The SDK requires Go 1.21 or higher.

## 📲 Installation

First time using Mercado Pago? Create your [Mercado Pago account](https://www.mercadopago.com), if you don't have one already.

1. Install the Mercado Pago SDK for Go:
```sh
$ go install github.com/mercadopago/sdk-go
```

2. Import the SDK into your Go code:

```go
import "github.com/mercadopago/sdk-go"
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
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{CARD_TOKEN}}",
		Installments: 1,
	}

	client := payment.NewClient(cfg)
	pay, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pay)
}
```

### SDK Configuration

Before making API requests, you need to initialize the SDK with your access token:

```go
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
```

### Making API Requests

To make requests to the Mercado Pago API, you can use the methods provided by the SDK. For example, to create a payment, you can use the `Payment.Create` method:

```go
	req := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{CARD_TOKEN}}",
		Installments: 1,
	}

	client := payment.NewClient(cfg)
```

### Exception throwing handling

In both cases, client configuration and payment creation, the variable `err` is available, which will contain any error thrown, it is important to handle these errors in the best possible way
```go
	if err != nil {
		fmt.Println(err)
		return
	}
```

For more details on the available methods and request parameters, please refer to the [Go Reference](https://pkg.go.dev/github.com/mercadopago/sdk-go) documentation.

## 📚 Documentation

See our documentation for more details.

- Mercado Pago API Reference: [English](https://www.mercadopago.com/developers/en/guides)
- Mercado Pago API Reference: [Spanish](https://www.mercadopago.com/developers/es/guides)

## 🤝 Contributing

All contributions are welcome, ranging from people wanting to triage issues, others wanting to write documentation, to people wanting to contribute code.

Please read and follow our [contribution guidelines](CONTRIBUTING.md). Contributions not following these guidelines will be disregarded. The guidelines are in place to make all of our lives easier and make contribution a consistent process for everyone.

### Patches to Version 1.x.x

Since the release of version 2.0.0, version 1 is deprecated and will not be receiving new features, only bug fixes. If you need to submit PRs for that version, please do so by using `develop-v1` as your base branch.

## ❤️ Support

If you require technical support, please contact our support team at our developers site: [English](https://www.mercadopago.com/developers/en/support/center/contact) / [Spanish](https://www.mercadopago.com/developers/es/support/center/contact)

## 🏻 License

```
MIT license. Copyright (c) 2024 - Mercado Pago / Mercado Libre 
For more information, see the LICENSE file.
```
