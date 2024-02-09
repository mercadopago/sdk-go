package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment/refund"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := refund.NewClient(cfg)

	var paymentID int64 = 12233344

	refund, err := client.List(context.Background(), paymentID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(refund)
}
