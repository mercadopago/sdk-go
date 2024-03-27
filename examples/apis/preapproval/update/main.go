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

	client := preapproval.NewClient(cfg)

	createRequest := preapproval.Request{
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

	resource, err := client.Create(context.Background(), createRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	updateRequest := preapproval.UpdateRequest{
		AutoRecurring: &preapproval.AutoRecurringUpdateRequest{
			TransactionAmount: 100,
		},
		BackURL:           "https://www.yoursite.com",
		ExternalReference: "Ref-123",
		Reason:            "Yoga Class",
	}

	resource, err = client.Update(context.Background(), resource.ID, updateRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
