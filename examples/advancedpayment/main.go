package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mercadopago/sdk-go/pkg/advancedpayment"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	cfg, err := config.New("<YOUR_ACCESS_TOKEN>")
	if err != nil {
		log.Fatal(err)
	}

	client := advancedpayment.NewClient(cfg)

	// Create an advanced (split) payment
	request := advancedpayment.Request{
		ApplicationID: "<YOUR_APPLICATION_ID>",
		Payments: []advancedpayment.PaymentRequest{
			{
				PaymentMethodID:   "master",
				PaymentTypeID:     "credit_card",
				Token:             "<CARD_TOKEN>",
				TransactionAmount: 100.0,
				Installments:      1,
				ProcessingMode:    "aggregator",
				Description:       "Split payment example",
			},
		},
		Disbursements: []advancedpayment.DisbursementRequest{
			{
				CollectorID:    488656838,
				Amount:         80.0,
				ApplicationFee: 2.0,
			},
		},
		Payer:             &advancedpayment.PayerRequest{Email: "buyer@example.com"},
		ExternalReference: "ADV-REF-001",
		Capture:           false,
	}

	result, err := client.Create(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Advanced Payment ID: %d, Status: %s\n", result.ID, result.Status)

	// Capture the payment
	captured, err := client.Capture(context.Background(), result.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Captured status: %s\n", captured.Status)
}
