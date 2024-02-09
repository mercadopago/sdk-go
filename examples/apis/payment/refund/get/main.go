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

	var paymentID int64 = 12344555
	var refundID int64 = 12344555

	refund, err := client.Get(context.Background(), paymentID, refundID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(refund)
}
