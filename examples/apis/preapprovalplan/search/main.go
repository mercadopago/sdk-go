package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"

	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preapprovalplan.NewClient(cfg)

	filters := preapprovalplan.SearchRequest{
		Limit:  10,
		Offset: 10,
		Filters: map[string]string{
			"status": "active",
		},
	}

	result, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, plan := range result.Results {
		fmt.Println(plan)
	}
}
