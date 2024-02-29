package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/invoice"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := invoice.NewClient(cfg)

	invoiceID := "123"

	inv, err := client.Get(context.Background(), invoiceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(inv)
}
