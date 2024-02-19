package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func main() {
	accessToken := "APP_USR-4849723703374061-053108-80867bffcb4a85cda0cd797f6c40cf28-1340175910"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := merchantorder.NewClient(cfg)

	filter := make(map[string]string)
	filter["preference_id"] = "1340175910-f2694bdc-7562-499e-a373-057cce3a027b"
	req := merchantorder.SearchRequest{
		Limit:   "10",
		Offset:  "0",
		Filters: filter,
	}

	search, err := client.Search(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, order := range search.Elements {
		fmt.Println(order)
	}
}
