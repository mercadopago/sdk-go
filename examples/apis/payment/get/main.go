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
	paymentRequest := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "pix",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
	}
	client := payment.NewClient(cfg)

	payment, err := client.Create(context.Background(), paymentRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err = client.Get(context.Background(), payment.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(payment)
}
