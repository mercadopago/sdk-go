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

	request := preference.Request{
		Items: []preference.PreferenceItemRequest{
			{
				ID:          "123",
				Title:       "Title",
				UnitPrice:   100,
				Quantity:    1,
				Description: "Description",
			},
		},
	}

	client := preference.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
