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

	client := point.NewClient(cfg)

	request := point.Request{
		Amount:      1500,
		Description: "your payment intent description",
		AdditionalInfo: &point.AdditionalInfoRequest{
			PrintOnTerminal:   false,
			ExternalReference: "4561ads-das4das4-das4754-das456",
		},
		Payment: &point.PaymentRequest{
			Installments:     1,
			Type:             "credit_card",
			InstallmentsCost: "seller",
		},
	}

	resource, err := client.Create(context.Background(), "{{DEVICE_ID}}", request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
