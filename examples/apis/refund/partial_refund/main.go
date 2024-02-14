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
	paymentRequest := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "master",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Token:        "{{TOKEN}}",
		Installments: 1,
		Capture:      false,
	}

	paymentClient := payment.NewClient(cfg)
	payment, err := paymentClient.Create(context.Background(), paymentRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	partialAmount := payment.TransactionAmount - 10.0

	refundClient := refund.NewClient(cfg)
	refund, err := refundClient.CreatePartialRefund(context.Background(), partialAmount, payment.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(refund)
}
