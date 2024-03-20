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

	client := merchantorder.NewClient(cfg)

	filters := make(map[string]string)
	filters["preference_id"] = "{{PREFERENCE_ID}}"
	request := merchantorder.SearchRequest{
		Filters: filters,
	}

	resource, err := client.Search(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resource.Elements {
		fmt.Println(v)
	}
}
