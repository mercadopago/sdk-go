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

	request := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "pix",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
	}

	client := payment.NewClient(cfg)
	pay, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	pay, err = client.Get(context.Background(), pay.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pay)
}
