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

	card, err := client.Delete(context.Background(), "{{CUSTOMER_ID}}", "{{CARD_ID}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(card)
}
