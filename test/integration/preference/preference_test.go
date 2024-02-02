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
		c, err := config.New(os.Getenv("at"))
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

		pmc := preference.NewClient(c)
		res, err := pmc.Create(context.Background(), dto)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_preference", func(t *testing.T) {
		c, err := config.New("TEST-4679935697572392-071411-a9722b82869609999cd91f0db60598f0-1273205088")
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

		pmc := preference.NewClient(c)
		res, _ := pmc.Create(context.Background(), dto)

		res, err = pmc.Get(context.Background(), res.ID)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_preference", func(t *testing.T) {
		c, err := config.New(os.Getenv("at"))
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
		pmc := preference.NewClient(c)
		res, err := pmc.Create(context.Background(), create)
		if res == nil {
			t.Error("res can't be nil")
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
					Quantity:    1,
					Description: "Description",
				},
			},
		}

		if res != nil {
			res, err = pmc.Update(context.Background(), res.ID, update)

			if res == nil {
				t.Error("res can't be nil")
			}
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_preference", func(t *testing.T) {
		c, err := config.New(os.Getenv("at"))
		if err != nil {
			t.Fatal(err)
		}

		filters := preference.SearchRequest{
			Limit:  10,
			Offset: 10,
		}

		pmc := preference.NewClient(c)
		res, err := pmc.Search(context.Background(), filters)

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
