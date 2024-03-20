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

	request := invoice.SearchRequest{
		Limit:  10,
		Offset: 10,
		Filters: map[string]string{
			"preapproval_id": "preapproval_id",
		},
	}

	resource, err := client.Search(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resource.Results {
		fmt.Println(v)
	}
}
