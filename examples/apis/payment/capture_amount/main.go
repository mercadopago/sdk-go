package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	at := "TEST-640110472259637-071923-a761f639c4eb1f0835ff7611f3248628-793910800"
	c, err := config.New(at)
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

	// Capture amount.
	result, err = client.CaptureAmount(context.Background(), result.ID, 100.1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
