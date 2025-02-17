package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	orderID := "{{ORDER_ID}}"
	transactionID := "{{TRANSACTION_ID}}"

	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)

	err = client.DeleteTransaction(context.Background(), orderID, transactionID)
	if err != nil {
		fmt.Println("Error in delete transaction:", err)
		return
	}

	fmt.Println("Deleted ")
}
