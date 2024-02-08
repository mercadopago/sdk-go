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
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		dto := preference.Request{
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

		client := preference.NewClient(c)
		res, err := client.Create(context.Background(), dto)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_preference", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		dto := preference.Request{
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

		client := preference.NewClient(c)
		res, _ := client.Create(context.Background(), dto)

		res, err = client.Get(context.Background(), res.ID)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_preference", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		create := preference.Request{
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
		client := preference.NewClient(c)
		res, err := client.Create(context.Background(), create)
		if res == nil {
			t.Error("res can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		update := preference.Request{
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

		res, err = client.Update(context.Background(), update, res.ID)

		if res == nil {
			t.Error("res can't be nil")
		}

		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_preference", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		filters := preference.SearchRequest{
			Limit:  "10",
			Offset: "10",
		}

		client := preference.NewClient(c)
		res, err := client.Search(context.Background(), filters)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
