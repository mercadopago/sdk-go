package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

func TestPreApprovalPlan(t *testing.T) {
	t.Run("should_create_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapprovalplan.NewClient(cfg)

		request := preapprovalplan.Request{
			AutoRecurring: &preapprovalplan.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: &preapprovalplan.PaymentMethodsAllowedRequest{
				PaymentTypes: []preapprovalplan.PaymentTypeRequest{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethodRequest{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil || resource.ID == "" {
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

		request := preapprovalplan.Request{
			AutoRecurring: &preapprovalplan.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: &preapprovalplan.PaymentMethodsAllowedRequest{
				PaymentTypes: []preapprovalplan.PaymentTypeRequest{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethodRequest{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		resource, err = client.Get(context.Background(), resource.ID)
		if resource == nil || resource.ID == "" {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_update_preapproval_plan", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := preapprovalplan.NewClient(cfg)

		request := preapprovalplan.Request{
			AutoRecurring: &preapprovalplan.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "days",
				TransactionAmount: 5,
				CurrencyID:        "BRL",
			},
			BackURL: "https://www.yoursite.com",
			PaymentMethodsAllowed: &preapprovalplan.PaymentMethodsAllowedRequest{
				PaymentTypes: []preapprovalplan.PaymentTypeRequest{
					{
						ID: "credit_card",
					},
				},
				PaymentMethods: []preapprovalplan.PaymentMethodRequest{
					{
						ID: "bolbradesco",
					},
				},
			},
			Reason: "Yoga classes",
		}

		resource, err := client.Create(context.Background(), request)
		if resource == nil {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		request = preapprovalplan.Request{
			AutoRecurring: &preapprovalplan.AutoRecurringRequest{
				Frequency:         1,
				FrequencyType:     "months",
				TransactionAmount: 10,
				BillingDay:        10,
				Repetitions:       12,
				CurrencyID:        "BRL",
			},
		}

		resource, err = client.Update(context.Background(), resource.ID, request)
		if resource == nil {
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
			Limit:  10,
			Offset: 10,
		}

		client := preapprovalplan.NewClient(cfg)
		resource, err := client.Search(context.Background(), filters)

		if resource == nil || resource.Results[0].ID == "" {
			t.Error("preapproval_plan can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
