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
	ppi, err := client.Search(context.Background(), "{{PAYMENT_INTENT_ID}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ppi)
}
