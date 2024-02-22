package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/point"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := point.Request{
		Amount:      1500,
		Description: "your payment intent description",
		AdditionalInfo: point.AdditionalInfo{
			PrintOnTerminal:   false,
			ExternalReference: "4561ads-das4das4-das4754-das456",
		},
		Payment: point.Payment{
			Installments:     1,
			Type:             "credit_card",
			InstallmentsCost: "seller",
		},
	}

	client := point.NewClient(cfg)
	ppi, err := client.Create(context.Background(), "{{DEVICE_ID}}", req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ppi)
}
