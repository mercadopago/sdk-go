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

	client := preapprovalplan.NewClient(cfg)
	result, err := client.Create(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
