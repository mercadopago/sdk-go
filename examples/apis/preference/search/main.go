package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preference.NewClient(cfg)

	filter := make(map[string]string)
	filter["external_reference"] = "wee3rffee23"
	filters := preference.SearchRequest{
		Limit:   10,
		Offset:  10,
		Filters: filter,
	}

	search, err := client.Search(context.Background(), filters)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pref := range search.Elements {
		fmt.Println(pref)
	}
}
