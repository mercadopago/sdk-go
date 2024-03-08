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
	req := payment.Request{
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
	pay, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	partialAmount := pay.TransactionAmount - 10.0

	refundClient := refund.NewClient(cfg)
	ref, err := refundClient.CreatePartialRefund(context.Background(), pay.ID, partialAmount)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ref)
}
