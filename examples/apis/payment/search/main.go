package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := payment.SearchRequest{
		Filters: map[string]string{
			"external_reference": "abc_def_ghi_123_456123",
		},
	}

	client := payment.NewClient(cfg)
	resource, err := client.Search(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resource.Results {
		fmt.Println(v)
	}
}
