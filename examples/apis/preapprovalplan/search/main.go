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
		Limit:  "10",
		Offset: "10",
	}

	search, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pref := range search.Results {
		fmt.Println(pref)
	}
}
