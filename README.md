# Mercado Pago SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mercadopago/dx-go.svg)](https://pkg.go.dev/github.com/mercadopago/dx-go)
[![License](https://img.shields.io/apm/l/vim-mode)](https://github.com/mercadopago/sdk-go)

A comprehensive Go client library for integrating with the Mercado Pago API.

## 💡 Requirements

The SDK requires Go 1.15 or higher.

## 📲 Installation

First time using Mercado Pago? Create your [Mercado Pago account](https://www.mercadopago.com), if you don't have one already.

1. Install the Mercado Pago SDK for Go:
```sh
$ go get github.com/mercadopago/dx-go
```

2. Import the SDK into your Go code:

```go
import "github.com/mercadopago/dx-go"
```

That's it! The Mercado Pago SDK for Go has been successfully installed.

## 🌟 Getting Started

Simple usage looks like:

```go
package main

import (
	"fmt"

	mp "github.com/mercadopago/dx-go"
)

func main() {
	mp.SDK.Init("YOUR_ACCESS_TOKEN")

	payment, err := mp.SDK.Payment.Create(&mp.PaymentCreateRequest{
		TransactionAmount: 1000.0,
		Token:             "your_cardtoken",
		Description:       "description",
		Installments:      1,
		PaymentMethodID:   "visa",
		Payer: &mp.PaymentPayerRequest{
			Email: "dummy_email",
		},
	})

	if err != nil {
		fmt.Printf("Mercado Pago Error. Status: %d, Content: %s\n", err.Status, err.Content)
		return
	}

	fmt.Println(payment)
}
```

### SDK Configuration

Before making API requests, you need to initialize the SDK with your access token. You can do this by calling the `Init` function:

```go
mp.SDK.Init("YOUR_ACCESS_TOKEN")
```

### Making API Requests

To make requests to the Mercado Pago API, you can use the methods provided by the SDK. For example, to create a payment, you can use the `Payment.Create` method:

```go
payment, err := mp.SDK.Payment.Create(&mp.PaymentCreateRequest{
	TransactionAmount: 1000.0,
	Token:             "your_cardtoken",
	Description:       "description",
	Installments:      1,
	PaymentMethodID:   "visa",
	Payer: &mp.PaymentPayerRequest{
		Email: "dummy_email",
	},
})

if err != nil {
	fmt.Printf("Mercado Pago Error. Status: %d, Content: %s\n", err.Status, err.Content)
	return
}

fmt.Println(payment)
```

For more details on the available methods and request parameters, please refer to the [Go Reference](https://pkg.go.dev/github.com/mercadopago/dx-go) documentation.

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
