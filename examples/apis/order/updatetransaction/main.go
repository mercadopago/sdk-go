package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{Access_token}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	orderRequest := order.Request{
		Type:              "online",
		ProcessingMode:    "manual",
		TotalAmount:       "100.00",
		ExternalReference: "ext_ref_1234",
		Payer: order.PayerRequest{
			Email: "{{EMAIL}}",
		},
	}

	orderResource, err := client.Create(context.Background(), orderRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	orderID := orderResource.ID
	fmt.Println("Order created with success. Order IDL: ", orderID)

	//create a transaction:
	transactionRequest := order.TransactionRequest{
		Payments: []order.PaymentRequest{
			{
				Amount: "100.00",
				PaymentMethod: order.PaymentMethodRequest{
					ID:           "master",
					Token:        "{{CARD_TOKEN}}",
					Type:         "credit_card",
					Installments: 1,
				},
			},
		},
	}
	transactionCreated, err := client.CreateTransaction(context.Background(), orderID, transactionRequest)
	if err != nil {
		fmt.Println("Error in create transaction", err)
		return
	}

	fmt.Println("Success in create a Transaction", transactionCreated)
}
