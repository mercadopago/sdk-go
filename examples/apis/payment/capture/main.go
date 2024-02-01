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

	// Create payment.
	dto := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{CARD_TOKEN}}",
		Installments: 1,
		Capture:      false,
	}

	client := payment.NewClient(cfg)
	result, err := client.Create(context.Background(), dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Capture.
	result, err = client.Capture(context.Background(), result.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
