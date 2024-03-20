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

	paymentClient := payment.NewClient(cfg)
	refundClient := refund.NewClient(cfg)

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

	paymentResource, err := paymentClient.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	partialAmount := paymentResource.TransactionAmount - 10.0

	refundResource, err := refundClient.CreatePartialRefund(context.Background(), paymentResource.ID, partialAmount)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(refundResource)
}
