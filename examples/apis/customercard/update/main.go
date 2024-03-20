package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/customercard"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := customercard.NewClient(cfg)

	request := customercard.Request{Token: "{{CARD_TOKEN}}"}

	resource, err := client.Update(context.Background(), "{{CUSTOMER_ID}}", "{{CARD_ID}}", request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
