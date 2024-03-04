package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/refund"
)

func TestRefund(t *testing.T) {
	t.Run("should_create_refund", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "visa",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
			// Need to get a token from a card.
			Token:        "",
			Installments: 1,
			Capture:      false,
		}

		paymentClient := payment.NewClient(cfg)
		pay, err := paymentClient.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		refundClient := refund.NewClient(cfg)
		ref, err := refundClient.Create(context.Background(), pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_create_partial_refund", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "visa",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
			// Need to get a token from a card.
			Token:        "",
			Installments: 1,
			Capture:      false,
		}

		paymentClient := payment.NewClient(cfg)
		pay, err := paymentClient.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		partialAmount := req.TransactionAmount - 5.0

		refundClient := refund.NewClient(cfg)
		ref, err := refundClient.CreatePartialRefund(context.Background(), partialAmount, pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_refund", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "visa",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
			// Need to get a token from a card.
			Token:        "",
			Installments: 1,
			Capture:      false,
		}

		paymentClient := payment.NewClient(cfg)
		pay, err := paymentClient.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		refundClient := refund.NewClient(cfg)
		ref, err := refundClient.Create(context.Background(), pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		ref, err = refundClient.Get(context.Background(), pay.ID, ref.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if ref.ID == 0 {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_list_refund", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "visa",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
			// Need to get a token from a card.
			Token:        "",
			Installments: 1,
			Capture:      false,
		}

		paymentClient := payment.NewClient(cfg)
		pay, err := paymentClient.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		partialAmount := req.TransactionAmount - 5.0

		// Partial refund
		refundClient := refund.NewClient(cfg)
		ref, err := refundClient.CreatePartialRefund(context.Background(), partialAmount, pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Total refund
		ref, err = refundClient.Create(context.Background(), pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		refunds, err := refundClient.List(context.Background(), pay.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(refunds) != 2 {
			t.Error("size can't be different of 2")
		}
	})
}
