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
		Type:              "{{TYPE}}",
		TotalAmount:       "{{TOTAL_AMOUNT}}",
		ExternalReference: "{{EXTERNAL_REFERENCE}}",
    Currency:          "BRL",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "{{AMOUNT}}",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "{{PAYMENT_METHOD_ID}}",
						Token:        "{{CARD_TOKEN}}",
						Type:         "{{TYPE}}",
						Installments: 1,
					},
				},
			},
		},
		Payer: &order.PayerRequest{
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
