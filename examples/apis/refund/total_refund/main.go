package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/refund"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create payment.
	request := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "master",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{TOKEN}}",
		Installments: 1,
		Capture:      false,
	}

	client := payment.NewClient(cfg)
	pay, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	refundClient := refund.NewClient(cfg)
	ref, err := refundClient.Create(context.Background(), pay.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ref)
}
