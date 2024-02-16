package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := payment.NewClient(cfg)
	var paymentID int64 = 123

	pay, err := client.Cancel(context.Background(), paymentID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pay)
}
