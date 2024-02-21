package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create preference.
	prefReq := preference.Request{
		ExternalReference: uuid.New().String(),
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
	pref, err := preferenceClient.Create(context.Background(), prefReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create merchant order.
	req := merchantorder.Request{
		ExternalReference: pref.ExternalReference,
		PreferenceID:      pref.ID,
		Collector: &merchantorder.CollectorRequest{
			ID: pref.CollectorID,
		},
		SiteID: pref.SiteID,
		Items: []merchantorder.ItemRequest{
			{
				CategoryID:  pref.Items[0].CategoryID,
				CurrencyID:  pref.Items[0].CurrencyID,
				Description: pref.Items[0].Description,
				PictureURL:  pref.Items[0].PictureURL,
				Title:       pref.Items[0].Title,
				Quantity:    pref.Items[0].Quantity,
				UnitPrice:   pref.Items[0].UnitPrice,
			},
		},
	}

	client := merchantorder.NewClient(cfg)
	order, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(order)
}
