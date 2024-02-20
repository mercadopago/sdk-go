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

	// Create payment.
	req := merchantorder.Request{
		ExternalReference: "default",
		PreferenceID:      "123456789",
		Collector: &merchantorder.CollectorRequest{
			ID: 123456789,
		},
		SiteID: "MLB",
		Items: []merchantorder.ItemRequest{
			{
				CategoryID:  "MLB123456789",
				CurrencyID:  "BRL",
				Description: "description",
				PictureURL:  "https://http2.mlstatic.com/D_NQ_NP_652451-MLB74602308021_022024-F.jpg",
				Title:       "title",
				Quantity:    1,
				UnitPrice:   1,
			},
		},
		ApplicationID: "123456789",
		Version:       1,
	}

	client := merchantorder.NewClient(cfg)
	order, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(order)
}
