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

		prefReq := preference.Request{
			Items: []preference.ItemRequest{
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
		if pref == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
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

		prefReq := preference.Request{
			Items: []preference.ItemRequest{
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
		if pref == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
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
		if order == nil || order.ID == 0 {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
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

		order, err = client.Update(context.Background(), order.ID, req)
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

		prefReq := preference.Request{
			Items: []preference.ItemRequest{
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
		if pref == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
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
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		order, err = client.Get(context.Background(), order.ID)
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

		req := merchantorder.SearchRequest{
			Limit: 5,
		}

		client := merchantorder.NewClient(cfg)
		order, err := client.Search(context.Background(), req)
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
