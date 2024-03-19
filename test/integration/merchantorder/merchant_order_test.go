package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func TestMerchantOrder(t *testing.T) {
	t.Run("should_create_merchant_order", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		preferenceClient := preference.NewClient(cfg)
		merchantOrderClient := merchantorder.NewClient(cfg)

		prefReq := preference.Request{
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

		preferenceResource, err := preferenceClient.Create(context.Background(), prefReq)
		if preferenceResource == nil {
			t.Error("preferenceResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		request := merchantorder.Request{
			ExternalReference: preferenceResource.ExternalReference,
			PreferenceID:      preferenceResource.ID,
			Collector: &merchantorder.CollectorRequest{
				ID: preferenceResource.CollectorID,
			},
			SiteID: preferenceResource.SiteID,
			Items: []merchantorder.ItemRequest{
				{
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

		merchantOrderResource, err := merchantOrderClient.Create(context.Background(), request)
		if merchantOrderResource == nil {
			t.Error("merchantOrderResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_merchant_order", func(t *testing.T) {
		cfg, err := config.New("TEST-4849723703374061-053108-98d6fdf742a963513320c567195b5cd6-1340175910")
		if err != nil {
			t.Fatal(err)
		}

		preferenceClient := preference.NewClient(cfg)
		merchantOrderClient := merchantorder.NewClient(cfg)

		prefReq := preference.Request{
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

		preferenceResource, err := preferenceClient.Create(context.Background(), prefReq)
		if preferenceResource == nil {
			t.Error("preferenceResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		createReq := merchantorder.Request{
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

		merchantOrderResource, err := merchantOrderClient.Create(context.Background(), createReq)
		if merchantOrderResource == nil || merchantOrderResource.ID == 0 {
			t.Error("merchantOrderResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Update merchant order.
		request := merchantorder.UpdateRequest{
			PreferenceID: preferenceResource.ID,
			SiteID:       preferenceResource.SiteID,
			Items: []merchantorder.ItemUpdateRequest{
				{
					ID:       merchantOrderResource.Items[0].ID,
					Quantity: 2,
				},
			},
		}

		merchantOrderResource, err = merchantOrderClient.Update(context.Background(), merchantOrderResource.ID, request)
		if merchantOrderResource == nil {
			fmt.Println(err)
			t.Error("merchantOrderResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_merchant_order", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		preferenceClient := preference.NewClient(cfg)
		merchantOrderClient := merchantorder.NewClient(cfg)

		prefReq := preference.Request{
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

		preferenceResource, err := preferenceClient.Create(context.Background(), prefReq)
		if preferenceResource == nil {
			t.Error("preferenceResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		request := merchantorder.Request{
			ExternalReference: preferenceResource.ExternalReference,
			PreferenceID:      preferenceResource.ID,
			Collector: &merchantorder.CollectorRequest{
				ID: preferenceResource.CollectorID,
			},
			SiteID: preferenceResource.SiteID,
			Items: []merchantorder.ItemRequest{
				{
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

		merchantOrderResource, err := merchantOrderClient.Create(context.Background(), request)
		if merchantOrderResource == nil {
			t.Error("merchantOrderResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		merchantOrderResource, err = merchantOrderClient.Get(context.Background(), merchantOrderResource.ID)
		if merchantOrderResource == nil {
			t.Error("merchantOrderResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
		if merchantOrderResource.ID == 0 {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_search_merchant_order", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := merchantorder.NewClient(cfg)

		request := merchantorder.SearchRequest{
			Limit: 5,
		}

		resource, err := client.Search(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
