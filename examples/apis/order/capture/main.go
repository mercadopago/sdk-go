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

	captureResp, err := client.Capture(context.Background(), orderID)
	if err != nil {
		fmt.Println("Error capturing the order:", err)
		return
	}

	if captureResp.Status == "processed" {
		fmt.Printf("Order %s captured successfully.\n", orderID)
	}
}
