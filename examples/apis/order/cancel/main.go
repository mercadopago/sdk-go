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

	cancelResp, err := client.Cancel(context.Background(), orderID)
	if err != nil {
		fmt.Println("Error canceling the order:", err)
		return
	}

	if cancelResp.Status == "canceled" {
		fmt.Printf("Order %s canceled successfully.\n", orderID)
	}
}
