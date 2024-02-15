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

	preferenceRequest := preference.Request{
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

	preferenceClient := preference.NewClient(cfg)
	preference, err := preferenceClient.Create(context.Background(), preferenceRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(preference)
}
