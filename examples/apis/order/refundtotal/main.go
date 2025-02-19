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

	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)

	_, err = client.Refund(context.Background(), orderID, nil)
	if err != nil {
		fmt.Println("Error in refund transaction:", err)
		return
	}

	fmt.Println("Refunded ")
}
