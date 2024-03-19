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

		resource, err := preferenceClient.Create(context.Background(), prefReq)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		request := merchantorder.Request{
			ExternalReference: resource.ExternalReference,
			PreferenceID:      resource.ID,
			Collector: &merchantorder.CollectorRequest{
				ID: resource.CollectorID,
			},
			SiteID: resource.SiteID,
			Items: []merchantorder.ItemRequest{
				{
					CategoryID:  resource.Items[0].CategoryID,
					CurrencyID:  resource.Items[0].CurrencyID,
					Description: resource.Items[0].Description,
					PictureURL:  resource.Items[0].PictureURL,
					Title:       resource.Items[0].Title,
					Quantity:    resource.Items[0].Quantity,
					UnitPrice:   resource.Items[0].UnitPrice,
				},
			},
		}

		order, err := merchantOrderClient.Create(context.Background(), request)
		if order == nil {
			t.Error("merchant order can't be nil")
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

		resource, err := preferenceClient.Create(context.Background(), prefReq)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		createReq := merchantorder.Request{
			ExternalReference: resource.ExternalReference,
			PreferenceID:      resource.ID,
			Collector: &merchantorder.CollectorRequest{
				ID: resource.CollectorID,
			},
			SiteID: resource.SiteID,
			Items: []merchantorder.ItemRequest{
				{
					ID:          resource.Items[0].ID,
					CategoryID:  resource.Items[0].CategoryID,
					CurrencyID:  resource.Items[0].CurrencyID,
					Description: resource.Items[0].Description,
					PictureURL:  resource.Items[0].PictureURL,
					Title:       resource.Items[0].Title,
					Quantity:    resource.Items[0].Quantity,
					UnitPrice:   resource.Items[0].UnitPrice,
				},
			},
		}

		order, err := merchantOrderClient.Create(context.Background(), createReq)
		if order == nil || order.ID == 0 {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Update merchant order.
		request := merchantorder.UpdateRequest{
			PreferenceID: resource.ID,
			SiteID:       resource.SiteID,
			Items: []merchantorder.ItemUpdateRequest{
				{
					ID:       order.Items[0].ID,
					Quantity: 2,
				},
			},
		}

		order, err = merchantOrderClient.Update(context.Background(), order.ID, request)
		if order == nil {
			fmt.Println(err)
			t.Error("merchant order can't be nil")
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

		resource, err := preferenceClient.Create(context.Background(), prefReq)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create merchant order.
		request := merchantorder.Request{
			ExternalReference: resource.ExternalReference,
			PreferenceID:      resource.ID,
			Collector: &merchantorder.CollectorRequest{
				ID: resource.CollectorID,
			},
			SiteID: resource.SiteID,
			Items: []merchantorder.ItemRequest{
				{
					CategoryID:  resource.Items[0].CategoryID,
					CurrencyID:  resource.Items[0].CurrencyID,
					Description: resource.Items[0].Description,
					PictureURL:  resource.Items[0].PictureURL,
					Title:       resource.Items[0].Title,
					Quantity:    resource.Items[0].Quantity,
					UnitPrice:   resource.Items[0].UnitPrice,
				},
			},
		}

		order, err := merchantOrderClient.Create(context.Background(), request)
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		order, err = merchantOrderClient.Get(context.Background(), order.ID)
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
		if order.ID == 0 {
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

		order, err := client.Search(context.Background(), request)
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
