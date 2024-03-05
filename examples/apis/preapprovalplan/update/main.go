package main

import (
	"context"
	"fmt"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"

	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	req := preapprovalplan.Request{
		AutoRecurring: preapprovalplan.AutoRecurringRequest{
			Frequency:         1,
			FrequencyType:     "days",
			TransactionAmount: 5,
			CurrencyID:        "BRL",
		},
		BackURL: "https://www.yoursite.com",
		PaymentMethodsAllowed: preapprovalplan.PaymentMethodsAllowedRequest{
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

	client := preapprovalplan.NewClient(cfg)

	pref, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	req = preapprovalplan.Request{
		AutoRecurring: preapprovalplan.AutoRecurringRequest{
			Frequency:         1,
			FrequencyType:     "months",
			TransactionAmount: 10,
			BillingDay:        10,
			Repetitions:       12,
			CurrencyID:        "BRL",
		},
	}

	pref, err = client.Update(context.Background(), req, pref.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pref)
}
