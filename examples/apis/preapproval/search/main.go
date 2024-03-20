package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preapproval.NewClient(cfg)

	filters := preapproval.SearchRequest{
		Limit:  10,
		Offset: 10,
		Filters: map[string]string{
			"payer_id": "123123123",
		},
	}

	resource, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resource.Results {
		fmt.Println(v)
	}
}
