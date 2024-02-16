package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/customer"
)

func TestCustomer(t *testing.T) {
	t.Run("should_create_customer", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := customer.NewClient(cfg)

		email := os.Getenv("TEST_EMAIL")
		req := customer.Request{Email: email}
		cus, err := client.Create(context.Background(), req)
		if cus == nil {
			t.Error("customer can't be nil")
			return
		}
		if cus.ID == "" {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_customer", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := customer.NewClient(cfg)

		email := os.Getenv("TEST_EMAIL")
		req := customer.SearchRequest{
			Filters: map[string]string{
				"email": email,
			},
		}
		res, err := client.Search(context.Background(), req)
		if res == nil {
			t.Error("customerSearch can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_customer", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := customer.NewClient(cfg)

		email := os.Getenv("TEST_EMAIL")
		req := customer.Request{Email: email}
		cus, err := client.Create(context.Background(), req)
		if cus == nil {
			t.Error("customer can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		res, err := client.Get(context.Background(), cus.ID)
		if res == nil {
			t.Error("customer can't be nil")
			return
		}
		if res.ID == "" {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_customer", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := customer.NewClient(cfg)

		email := os.Getenv("TEST_EMAIL")
		req := customer.Request{Email: email}
		cus, err := client.Create(context.Background(), req)
		if cus == nil {
			t.Error("customer can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		uReq := customer.Request{Description: "Description updated."}
		res, err := client.Update(context.Background(), cus.ID, uReq)
		if res == nil {
			t.Error("customer can't be nil")
			return
		}
		if res.Description != "Description updated." {
			t.Error("update failed")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
