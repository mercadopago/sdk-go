package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func TestMerchantOrder(t *testing.T) {
	t.Run("should_create_merchant_order", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

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
		if order == nil {
			t.Error("merchant order can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_merchant_order", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		req := merchantorder.UpdateRequest{
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
		order, err := client.Update(context.Background(), req, 123456789)
		if order == nil {
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

		client := merchantorder.NewClient(cfg)
		order, err := client.Get(context.Background(), 123456789)
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
			Limit: "5",
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
