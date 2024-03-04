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
	accessToken := "TEST-4718610619866357-092020-f30ef41ea2a9e7ad0fa7bc101b5508af-751574177"

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
				ID:          uuid.New().String(),
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
	createReq := merchantorder.Request{
		ExternalReference: pref.ExternalReference,
		PreferenceID:      pref.ID,
		Collector: &merchantorder.CollectorRequest{
			ID: pref.CollectorID,
		},
		SiteID: pref.SiteID,
		Items: []merchantorder.ItemRequest{
			{
				ID:          pref.Items[0].ID,
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
	order, err := client.Create(context.Background(), createReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update merchant order.
	req := merchantorder.UpdateRequest{
		PreferenceID: pref.ID,
		SiteID:       pref.SiteID,
		Items: []merchantorder.ItemUpdateRequest{
			{
				ID:       order.Items[0].ID,
				Quantity: 2,
			},
		},
	}

	// time.Sleep(time.Second * 10)

	order, err = client.Update(context.Background(), req, order.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(order)
}
