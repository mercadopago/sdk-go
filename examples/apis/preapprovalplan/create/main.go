package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

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

	client := preapprovalplan.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
