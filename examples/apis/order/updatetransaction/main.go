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
	orderID := "{{EXISTING_ORDER_ID}}"
	transactionID := "{{EXISTING_TRANSACTION_ID}}"

	// // Updating the transaction to change the number of installments
	updateRequest := order.PaymentRequest{
		PaymentMethod: &order.PaymentMethodRequest{Installments: 12},
	}

	ctx := context.Background()
	updatedTransaction, err := client.UpdateTransaction(ctx, orderID, transactionID, updateRequest)
	if err != nil {
		fmt.Printf("Failed to update transaction: %v\n", err)
		return
	}

	if updatedTransaction != nil {
		fmt.Printf("Transaction updated successfully! New installments: %v\n", updatedTransaction.PaymentMethod.Installments)
	}
}
