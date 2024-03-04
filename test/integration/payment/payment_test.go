package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
)

func TestPayment(t *testing.T) {
	t.Run("should_create_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(cfg)

		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		pay, err := client.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		req := payment.SearchRequest{
			Filters: map[string]string{
				"external_reference": "abc_def_ghi_123_456123",
			},
		}

		client := payment.NewClient(cfg)
		paymentSearch, err := client.Search(context.Background(), req)
		if paymentSearch == nil {
			t.Error("paymentSearch can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(cfg)
		paymentRequest := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		pay, err := client.Create(context.Background(), paymentRequest)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		pay, err = client.Get(context.Background(), pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_cancel_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(cfg)
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		pay, err := client.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		pay, err = client.Cancel(context.Background(), pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// We should validate how to test capture and capture amount.
	t.Run("should_capture_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(cfg)

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

		pay, err := client.Create(context.Background(), req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		pay, err = client.Capture(context.Background(), pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_capture_amount_payment", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(cfg)
		pay, err := client.CaptureAmount(context.Background(), 123, 100.1)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
