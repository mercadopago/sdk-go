# Mercado Pago SDK for Go

## Work In Progress
SDK development is still ongoing, we aim to release a reliable 1.0.0 soon.

[![Go Reference](https://pkg.go.dev/badge/github.com/mercadopago/sdk-go.svg)](https://pkg.go.dev/github.com/mercadopago/sdk-go)

![mercado-pago-image-7130x2250](https://github.com/mercadopago/sdk-go/assets/84413927/c18102b2-b4ed-46c9-9a83-b5e6a30d659b)

## Overview

A comprehensive Go client library for integrating with the Mercado Pago API.

## 💡 Requirements

The SDK requires Go 1.x.x or higher.

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
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(cfg)

	request := payment.Request{
		TransactionAmount: 105.1,
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{CARD_TOKEN}}",
		Installments: 1,
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

### Exception throwing handling

Every package methods returns two variables: response (type of the package) and error (type of the std lib), which will contain any error thrown. It is important to handle these errors in the best possible way.
```go
	resources, err := client.List(context.Background())
	if err != nil {
		// appropriate treatment
	}
```

For more details on the available methods and request parameters, please refer to the [Go Reference](https://pkg.go.dev/github.com/mercadopago/sdk-go) documentation.

## 📚 Documentation

See our documentation for more details.

- Mercado Pago API Reference: [English](https://www.mercadopago.com/developers/en/guides) | [Portuguese](https://www.mercadopago.com/developers/pt/guides) | [Spanish](https://www.mercadopago.com/developers/es/guides)

## 🤝 Contributing

All contributions are welcome, ranging from people wanting to triage issues, others wanting to write documentation, to people wanting to contribute with code.

Please read and follow our [contribution guidelines](CONTRIBUTING.md). Contributions not following these guidelines will be disregarded. The guidelines are in place to make all of our lives easier and make contribution a consistent process for everyone.

## ❤️ Support

If you require technical support, please contact our support team at our developers site: [English](https://www.mercadopago.com/developers/en/support/center/contact) | [Portuguese](https://www.mercadopago.com/developers/pt/support/center/contact) | [Spanish](https://www.mercadopago.com/developers/es/support/center/contact)
