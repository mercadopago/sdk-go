package integration

import (
	"context"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
	"os"
	"testing"
)

func TestPreApprovalPlan(t *testing.T) {
	t.Run("should_create_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapprovalplan.NewClient(cfg)

		req := preapprovalplan.Request{
			AutoRecurring: preapprovalplan.AutoRecurring{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: preapprovalplan.PaymentMethodsAllowed{
				PaymentTypes: []preapprovalplan.PaymentType{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethod{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		result, err := client.Create(context.Background(), req)
		if result == nil {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapprovalplan.NewClient(cfg)

		req := preapprovalplan.Request{
			AutoRecurring: preapprovalplan.AutoRecurring{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: preapprovalplan.PaymentMethodsAllowed{
				PaymentTypes: []preapprovalplan.PaymentType{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethod{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		result, err := client.Create(context.Background(), req)
		if result == nil {
			t.Error("preapproval_plan can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		result, err = client.Get(context.Background(), result.ID)
		if result == nil {
			t.Error("preapproval_plan can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}
		if result.ID == "" {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_update_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapprovalplan.NewClient(cfg)

		req := preapprovalplan.Request{
			AutoRecurring: preapprovalplan.AutoRecurring{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: preapprovalplan.PaymentMethodsAllowed{
				PaymentTypes: []preapprovalplan.PaymentType{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethod{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		result, err := client.Create(context.Background(), req)
		if result == nil {
			t.Error("preapproval_plan can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		req = preapprovalplan.Request{
			AutoRecurring: preapprovalplan.AutoRecurring{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 10,
				BillingDay:        10,
				Repetitions:       12,
				CurrencyID:        "BRL",
			},
		}

		result, err = client.Update(context.Background(), req, result.ID)
		if result == nil {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		filters := preapprovalplan.SearchRequest{
			Limit:  "10",
			Offset: "10",
		}

		client := preapprovalplan.NewClient(cfg)
		result, err := client.Search(context.Background(), filters)

		if result == nil || result.Results[0].ID == "" {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
