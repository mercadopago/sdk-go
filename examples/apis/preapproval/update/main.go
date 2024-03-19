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

	client := preapproval.NewClient(cfg)
	result, err := client.Create(context.Background(), req)
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

	result, err = client.Update(context.Background(), result.ID, update)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
