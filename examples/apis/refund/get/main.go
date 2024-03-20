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

	paymentID := 12344555
	refundID := 12344555

	resource, err := refundClient.Get(context.Background(), paymentID, refundID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
