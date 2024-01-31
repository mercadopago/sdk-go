package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create payment.
	dto := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
		},
		Token:        "cdec5028665c41976be212a7981437d6",
		Installments: 1,
		Capture:      false,
	}

	client := payment.NewClient(c)
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
