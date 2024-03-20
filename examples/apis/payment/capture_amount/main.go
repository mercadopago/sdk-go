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

	// Create payment.
	request := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{CARD_TOKEN}}",
		Installments: 1,
		Capture:      false,
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Capture amount.
	resource, err = client.CaptureAmount(context.Background(), resource.ID, 100.1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
