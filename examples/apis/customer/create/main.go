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

	email := "{{EMAIL}}"
	req := customer.Request{Email: email}

	client := customer.NewClient(cfg)
	res, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
