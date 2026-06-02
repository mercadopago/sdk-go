package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/disbursementrefund"
)

func main() {
	cfg, err := config.New("<YOUR_ACCESS_TOKEN>")
	if err != nil {
		log.Fatal(err)
	}

	client := disbursementrefund.NewClient(cfg)
	advancedPaymentID := 20458724
	disbursementID := 123456

	// List all refunds for an advanced payment
	refunds, err := client.ListAll(context.Background(), advancedPaymentID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Refunds count: %d\n", len(refunds))

	// Refund a specific disbursement by amount
	refund, err := client.Create(context.Background(), advancedPaymentID, disbursementID, disbursementrefund.Request{Amount: 50.0})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Refund ID: %d, Status: %s\n", refund.ID, refund.Status)
}
