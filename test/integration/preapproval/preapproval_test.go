package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

func TestPreApproval(t *testing.T) {
	t.Run("should_create_preapproval", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapproval.NewClient(cfg)

		request := preapproval.Request{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 100,
				CurrencyID:        "BRL",
			},
			BackURL:           "https://www.yoursite.com",
			ExternalReference: "Ref-123",
			PayerEmail:        "test_user_28355466@testuser.com",
			Reason:            "Yoga Class",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil || resource.ID == "" {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_preapproval", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapproval.NewClient(cfg)

		request := preapproval.Request{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 100,
				CurrencyID:        "BRL",
			},
			BackURL:           "https://www.yoursite.com",
			ExternalReference: "Ref-123",
			PayerEmail:        "test_user_28355466@testuser.com",
			Reason:            "Yoga Class",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		resource, err = client.Get(context.Background(), resource.ID)
		if resource == nil || resource.ID == "" {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_preapproval", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapproval.NewClient(cfg)

		request := preapproval.Request{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 100,
				CurrencyID:        "BRL",
			},
			BackURL:           "https://www.yoursite.com",
			ExternalReference: "Ref-123",
			PayerEmail:        "test_user_28355466@testuser.com",
			Reason:            "Yoga Class",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		updateRequest := preapproval.UpdateRequest{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 50.0,
				CurrencyID:        "BRL",
			},
		}

		resource, err = client.Update(context.Background(), resource.ID, updateRequest)
		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_preapproval", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapproval.NewClient(cfg)

		request := preapproval.Request{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 100,
				CurrencyID:        "BRL",
			},
			BackURL:           "https://www.yoursite.com",
			ExternalReference: "Ref-123",
			PayerEmail:        "test_user_28355466@testuser.com",
			Reason:            "Yoga Class",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		filters := preapproval.SearchRequest{
			Limit:  10,
			Offset: 10,
		}

		searchResource, err := client.Search(context.Background(), filters)
		if searchResource == nil {
			t.Error("searchResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
