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
	dto := payment.Request{
		TransactionAmount: 105.1,
		PaymentMethodID:   "pix",
		Payer: &payment.PayerRequest{
			Email: "{{EMAIL}}",
		},
	}

	client := payment.NewClient(cfg)

	result, err := client.Create(context.Background(), dto)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err = client.Get(context.Background(), result.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
