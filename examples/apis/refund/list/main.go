package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/refund"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	refundClient := refund.NewClient(cfg)

	paymentID := 12233344

	refunds, err := refundClient.List(context.Background(), paymentID)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ref := range refunds {
		fmt.Println(ref)
	}
}
