package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "online",
		TotalAmount:       "1000.00",
		ExternalReference: "ext_ref_1234",
		Transactions: order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "1000.00",
					PaymentMethod: order.PaymentMethodRequest{
						ID:           "master",
						Token:        "{{CARD_TOKEN}}",
						Type:         "credit_card",
						Installments: 1,
					},
				},
			},
		},
		Payer: order.PayerRequest{
			Email: "{{EMAIL}}",
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resource)
}
