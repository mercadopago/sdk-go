package integration

import (
	"context"
	"fmt"
	"math/rand"
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

		request := customer.Request{Email: generateEmail()}
		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
			return
		}
		if resource.ID == "" {
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

		request := customer.SearchRequest{
			Filters: map[string]string{
				"email": generateEmail(),
			},
		}
		resource, err := client.Search(context.Background(), request)
		if resource == nil {
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

		request := customer.Request{Email: generateEmail()}
		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		resource, err = client.Get(context.Background(), resource.ID)
		if resource == nil {
			t.Error("resource can't be nil")
			return
		}
		if resource.ID == "" {
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

		request := customer.Request{Email: generateEmail()}
		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		uReq := customer.Request{Description: "Description updated."}
		resource, err = client.Update(context.Background(), resource.ID, uReq)
		if resource == nil {
			t.Error("customer can't be nil")
			return
		}
		if resource.Description != "Description updated." {
			t.Error("update failed")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func generateEmail() string {
	id := rand.Intn(200000 /*max*/ -100000 /*min*/) + 100000 /*min*/
	return fmt.Sprintf("test_user_sdk_%d@testuser.com", id)
}
