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

	req := customercard.Request{Token: "{{CARD_TOKEN}}"}

	client := customercard.NewClient(cfg)
	card, err := client.Update(context.Background(), "{{CUSTOMER_ID}}", "{{CARD_ID}}", req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(card)
}
