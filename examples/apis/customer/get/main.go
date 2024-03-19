package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/customer"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := customer.NewClient(cfg)

	customerID := "{{CUSTOMER_ID}}"

	resource, err := client.Get(context.Background(), customerID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
