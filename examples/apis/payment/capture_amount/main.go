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
	paymentRequest := payment.Request{
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
	pay, err := client.Create(context.Background(), paymentRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Capture amount.
	result, err := client.CaptureAmount(context.Background(), pay.ID, 100.1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
