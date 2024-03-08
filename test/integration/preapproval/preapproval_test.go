package integration

import (
	"context"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
	"os"
	"testing"
)

func TestPreApproval(t *testing.T) {
	t.Run("should_create_preapproval", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapproval.NewClient(cfg)

		req := preapproval.Request{
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

		result, err := client.Create(context.Background(), req)
		if result == nil || result.ID == "" {
			t.Error("preapproval can't be nil")
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

		req := preapproval.Request{
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

		result, err := client.Create(context.Background(), req)
		if result == nil {
			t.Error("preapproval can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		result, err = client.Get(context.Background(), result.ID)
		if result == nil || result.ID == "" {
			t.Error("preapproval can't be nil")
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

		req := preapproval.Request{
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

		result, err := client.Create(context.Background(), req)
		if result == nil {
			t.Error("preapproval can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		update := preapproval.UpdateRequest{
			AutoRecurring: &preapproval.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 50.0,
				CurrencyID:        "BRL",
			},
		}

		result, err = client.Update(context.Background(), update, result.ID)
		if result == nil {
			t.Error("preapproval can't be nil")
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

		filters := preapproval.SearchRequest{
			Limit:  10,
			Offset: 10,
		}

		client := preapproval.NewClient(cfg)
		result, err := client.Search(context.Background(), filters)

		if result == nil || result.Results[0].ID == "" {
			t.Error("preapproval can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
