package main

// Mercado Pago Create Order — Automatic Payments (recurring charges).
//
// Demonstrates the two-step Automatic Payments flow:
//   1. First payment  — CVV-validated charge that registers the card credential.
//   2. Recurring charge — subsequent MIT charge without CVV, referencing step 1.
//
// Prerequisites:
//   - A customer created via POST /v1/customers                               → CUSTOMER_ID
//   - A payment profile created via POST /v1/customers/{id}/payment-profiles  → PAYMENT_PROFILE_ID
//
// See: https://www.mercadopago.com/developers/en/docs/automatic-payments-orders/overview

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/order"
)

func main() {
	accessToken    := "{{ACCESS_TOKEN}}"
	customerID     := "{{CUSTOMER_ID}}"
	profileID      := "{{PAYMENT_PROFILE_ID}}"
	payerEmail     := "{{PAYER_EMAIL}}"
	cardToken      := "{{CARD_TOKEN}}"

	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := order.NewClient(c)
	ctx    := context.Background()

	// ── Step 1: First payment ─────────────────────────────────────────────────
	// Registers the card credential with FirstPayment: true.
	// No PrevTransactionRef is needed on the first charge.
	firstPaymentReq := order.Request{
		Type:              "online",
		ProcessingMode:    "automatic",
		TotalAmount:       "100.00",
		ExternalReference: "subscription-001-payment-1",
		Payer: &order.PayerRequest{
			Email:      payerEmail,
			CustomerID: customerID,
		},
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "100.00",
					PaymentMethod: &order.PaymentMethodRequest{
						ID:           "master",
						Type:         "credit_card",
						Token:        cardToken,
						Installments: 1,
					},
					AutomaticPayments: &order.AutomaticPaymentsRequest{
						PaymentProfileID: profileID,
					},
					StoredCredential: &order.StoredCredentialRequest{
						PaymentInitiator: "customer",
						Reason:           "recurring",
						FirstPayment:     true,
					},
				},
			},
		},
	}

	firstOrder, err := client.Create(ctx, firstPaymentReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("First payment order ID: %s\n", firstOrder.ID)
	fmt.Printf("Status: %s\n", firstOrder.Status)

	// Save the payment ID — required as PrevTransactionRef in subsequent charges.
	if len(firstOrder.Transactions.Payments) == 0 {
		fmt.Println("No payments in response.")
		return
	}
	firstPaymentID := firstOrder.Transactions.Payments[0].ID
	fmt.Printf("First payment ID (save for next charge): %s\n", firstPaymentID)

	// ── Step 2: Recurring charge ──────────────────────────────────────────────
	// Subsequent MIT charge — no card token needed, uses the payment profile.
	// PrevTransactionRef links this charge to the original card-network authorization.
	recurringReq := order.Request{
		Type:              "online",
		ProcessingMode:    "automatic_async",
		TotalAmount:       "100.00",
		ExternalReference: "subscription-001-payment-2",
		Payer: &order.PayerRequest{
			Email:      payerEmail,
			CustomerID: customerID,
		},
		Transactions: &order.TransactionRequest{
			Payments: []order.PaymentRequest{
				{
					Amount: "100.00",
					AutomaticPayments: &order.AutomaticPaymentsRequest{
						PaymentProfileID: profileID,
						Retries:          3,
						ScheduleDate:     "2026-09-01T00:00:00.000-04:00",
						DueDate:          "2026-09-05T00:00:00.000-04:00",
					},
					StoredCredential: &order.StoredCredentialRequest{
						PaymentInitiator:   "merchant",
						Reason:             "recurring",
						FirstPayment:       false,
						PrevTransactionRef: firstPaymentID,
					},
					SubscriptionData: &order.SubscriptionDataRequest{
						InvoiceID:   "INV-002",
						BillingDate: "2026-08-01",
						SubscriptionSequence: &order.SubscriptionSequenceRequest{
							Number: 2,
							Total:  12,
						},
						InvoicePeriod: &order.InvoicePeriodRequest{
							Type:   "monthly",
							Period: 1,
						},
					},
				},
			},
		},
	}

	recurringOrder, err := client.Create(ctx, recurringReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nRecurring charge order ID: %s\n", recurringOrder.ID)
	fmt.Printf("Status: %s\n", recurringOrder.Status)
	fmt.Printf("Status detail: %s\n", recurringOrder.StatusDetail)
}
