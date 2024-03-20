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

	preferenceClient := preference.NewClient(cfg)
	client := merchantorder.NewClient(cfg)

	// Create preference.
	preferenceRequest := preference.Request{
		ExternalReference: uuid.New().String(),
		Items: []preference.PreferenceItemRequest{
			{
				ID:          uuid.New().String(),
				Title:       "Title",
				UnitPrice:   100,
				Quantity:    1,
				Description: "Description",
			},
		},
	}

	preferenceResource, err := preferenceClient.Create(context.Background(), preferenceRequest)
	if err != nil {
		fmt.Println(preferenceResource)
		return
	}

	// Create merchant order.
	merchantOrderRequest := merchantorder.Request{
		ExternalReference: preferenceResource.ExternalReference,
		PreferenceID:      preferenceResource.ID,
		Collector: &merchantorder.CollectorRequest{
			ID: preferenceResource.CollectorID,
		},
		SiteID: preferenceResource.SiteID,
		Items: []merchantorder.ItemRequest{
			{
				ID:          preferenceResource.Items[0].ID,
				CategoryID:  preferenceResource.Items[0].CategoryID,
				CurrencyID:  preferenceResource.Items[0].CurrencyID,
				Description: preferenceResource.Items[0].Description,
				PictureURL:  preferenceResource.Items[0].PictureURL,
				Title:       preferenceResource.Items[0].Title,
				Quantity:    preferenceResource.Items[0].Quantity,
				UnitPrice:   preferenceResource.Items[0].UnitPrice,
			},
		},
	}

	merchantOrderResource, err := client.Create(context.Background(), merchantOrderRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update merchant order.
	merchantOrderUpdateRequest := merchantorder.UpdateRequest{
		PreferenceID: preferenceResource.ID,
		SiteID:       preferenceResource.SiteID,
		Items: []merchantorder.ItemUpdateRequest{
			{
				ID:       merchantOrderResource.Items[0].ID,
				Quantity: 2,
			},
		},
	}

	merchantOrderResource, err = client.Update(context.Background(), merchantOrderResource.ID, merchantOrderUpdateRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(merchantOrderResource)
}
