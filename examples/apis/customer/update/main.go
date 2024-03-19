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
	request := customer.Request{Description: "Description updated."}

	resource, err := client.Update(context.Background(), customerID, request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
