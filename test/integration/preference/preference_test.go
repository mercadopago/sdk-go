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

		client := preference.NewClient(cfg)
		preference, err := client.Create(context.Background(), preferenceRequest)

		if preference == nil {
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

		client := preference.NewClient(cfg)
		preference, err := client.Create(context.Background(), preferenceRequest)
		if preference == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		preference, err = client.Get(context.Background(), preference.ID)

		if preference == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		if preference.ID == "" {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_update_preference", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
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
		client := preference.NewClient(cfg)
		preferenceCreated, err := client.Create(context.Background(), preferenceRequest)
		if preferenceCreated == nil {
			t.Error("preference can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		preferenceRequest = preference.Request{
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

		preference, err := client.Update(context.Background(), preferenceRequest, preferenceCreated.ID)

		if preference == nil {
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
			Limit:  "10",
			Offset: "10",
		}

		client := preference.NewClient(cfg)
		preference, err := client.Search(context.Background(), filters)

		if preference == nil {
			t.Error("preference can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
