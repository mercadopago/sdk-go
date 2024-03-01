package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/customer"
)

func main() {
	accessToken := "TEST-4718610619866357-092020-f30ef41ea2a9e7ad0fa7bc101b5508af-751574177"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	req := customer.SearchRequest{
		Filters: map[string]string{
			"EMAIL": "akdsokdoakdoasakdoas@testuser.com",
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
