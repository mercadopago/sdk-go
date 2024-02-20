package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	filter := make(map[string]string)
	filter["preference_id"] = "{{PREFERENCE_ID}}"
	req := merchantorder.SearchRequest{
		Filters: filter,
	}

	client := merchantorder.NewClient(cfg)
	search, err := client.Search(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, order := range search.Elements {
		fmt.Println(order)
	}
}
