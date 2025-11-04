package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	request := order.Request{
		Type:              "{{TYPE}}",
		TotalAmount:       "{{TOTAL_AMOUNT}}",
		ExternalReference: "{{EXTERNAL_REFERENCE}}",
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "{{AMOUNT}}",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "{{PAYMENT_METHOD_ID}}",
						Token:        "{{CARD_TOKEN}}",
						Type:         "{{TYPE}}",
						Installments: 1,
					},
				},
			},
		},
		Payer: &order.PayerRequest{
			Email: "{{EMAIL}}",
		},
		Config: &order.ConfigRequest{
			Online: &order.OnlineConfigRequest{
				TransactionSecurity: &order.TransactionSecurityRequest{
					Validation:     "always",
					LiabilityShift: "preferred",
				},
			},
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("== CRIAR ORDER (CHAMADA REAL API) ==")
	fmt.Println("✓ ORDER CRIADA COM SUCESSO!")
	fmt.Printf("ID: %s\n", resource.ID)
	fmt.Printf("Type: %s\n", resource.Type)
	fmt.Printf("External Reference: %s\n", resource.ExternalReference)
	fmt.Printf("Country: %s\n", resource.CountryCode)
	fmt.Printf("Status: %s (%s)\n", resource.Status, resource.StatusDetail)
	fmt.Printf("Capture Mode: %s\n", resource.CaptureMode)
	fmt.Printf("Processing Mode: %s\n", resource.ProcessingMode)
	fmt.Printf("Total Amount: %s | Paid: %s\n", resource.TotalAmount, resource.TotalPaidAmount)
	fmt.Printf("User ID (seller): %s\n", resource.UserID)
	fmt.Printf("Created: %s | Updated: %s\n", resource.CreatedDate, resource.LastUpdatedDate)
	fmt.Printf("Application ID: %s\n", resource.IntegrationData.ApplicationID)

	if len(resource.Transactions.Payments) > 0 {
		p := resource.Transactions.Payments[0]
		fmt.Println("\n-- Pagamento --")
		fmt.Printf("Payment ID: %s\n", p.ID)
		fmt.Printf("Reference: %s\n", p.ReferenceID)
		fmt.Printf("Status: %s (%s)\n", p.Status, p.StatusDetail)
		fmt.Printf("Amount: %s | Paid: %s | Attempts: %d\n", p.Amount, p.PaidAmount, p.AttemptNumber)
		fmt.Printf("Method: %s (%s) | Installments: %d\n", p.PaymentMethod.ID, p.PaymentMethod.Type, p.PaymentMethod.Installments)
	}

	fmt.Println("\n-- Resposta JSON Completa --")
	pretty, _ := json.MarshalIndent(resource, "", "  ")
	fmt.Println(string(pretty))
}
