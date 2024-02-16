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

	req := customer.SearchRequest{
		Filters: map[string]string{
			"email": "{{EMAIL}}",
		},
	}

	client := customer.NewClient(cfg)
	cus, err := client.Search(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cus)
}
