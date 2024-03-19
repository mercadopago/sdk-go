package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

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

	client := preapproval.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	update := preapproval.UpdateRequest{
		AutoRecurring: &preapproval.AutoRecurringRequest{
			Frequency:         1,
			FrequencyType:     "months",
			TransactionAmount: 100,
			CurrencyID:        "BRL",
		},
		BackURL:           "https://www.yoursite.com",
		ExternalReference: "Ref-123",
		Reason:            "Yoga Class",
	}

	resource, err = client.Update(context.Background(), resource.ID, update)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
