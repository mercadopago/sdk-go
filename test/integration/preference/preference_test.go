package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func TestPreference(t *testing.T) {
	t.Run("should_create_preference", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preference.NewClient(cfg)

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

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("preference can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_preference", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preference.NewClient(cfg)

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

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		resource, err = client.Get(context.Background(), resource.ID)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
		if resource.ID == "" {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_update_preference", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preference.NewClient(cfg)

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

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		request = preference.Request{
			Items: []preference.PreferenceItemRequest{
				{
					ID:          "123",
					Title:       "Title",
					UnitPrice:   100,
					Quantity:    2,
					Description: "Description",
				},
			},
		}

		resource, err = client.Update(context.Background(), resource.ID, request)
		if resource == nil {
			t.Error("preference can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_preference", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		filters := preference.SearchRequest{
			Limit:  10,
			Offset: 10,
		}

		client := preference.NewClient(cfg)
		resource, err := client.Search(context.Background(), filters)

		if resource == nil {
			t.Error("preference can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
