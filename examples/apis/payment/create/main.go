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

	dto := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "visa",
		Payer: &payment.PayerRequest{
			Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
		},
		Token:        "cdec5028665c41976be212a7981437d6",
		Installments: 1,
	}

	client := payment.NewClient(c)
	result, err := client.Create(context.Background(), dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
