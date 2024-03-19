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

	req := invoice.SearchRequest{
		Limit:  10,
		Offset: 10,
		Filters: map[string]string{
			"preapproval_id": "preapproval_id",
		},
	}

	result, err := client.Search(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, inv := range result.Results {
		fmt.Println(inv)
	}
}
