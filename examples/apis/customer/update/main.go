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

	customerID := "{{CUSTOMER_ID}}"
	req := customer.Request{Description: "Description updated."}

	client := customer.NewClient(cfg)
	cus, err := client.Update(context.Background(), customerID, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cus)
}
