package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/test"
)

var (
	cfg             = test.Config()
	paymentClient   = payment.NewClient(cfg)
	cardTokenClient = cardtoken.NewClient(cfg)
)

func TestPayment(t *testing.T) {
	t.Run("should_create_payment", func(t *testing.T) {
		ctx := context.Background()

		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New()),
			},
		}

		pay, err := paymentClient.Create(ctx, req)
		if pay == nil || pay.ID == 0 {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_payment", func(t *testing.T) {
		ctx := context.Background()

		req := payment.SearchRequest{
			Filters: map[string]string{
				"external_reference": "abc_def_ghi_123_456123",
			},
		}

		result, err := paymentClient.Search(ctx, req)
		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_payment", func(t *testing.T) {
		ctx := context.Background()

		// Create payment.
		paymentRequest := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("test_user_%s@testuser.com", uuid.New()),
			},
		}

		pay, err := paymentClient.Create(ctx, paymentRequest)
		if pay == nil || pay.ID == 0 {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Get payment.
		pay, err = paymentClient.Get(ctx, pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_cancel_payment", func(t *testing.T) {
		ctx := context.Background()

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
		}

		pay, err := paymentClient.Create(ctx, req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Cancel payment.
		pay, err = paymentClient.Cancel(ctx, pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay != nil && pay.Status != "cancelled" {
			t.Error("payment should be cancelled, but is wasn't")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_capture_payment", func(t *testing.T) {
		ctx := context.Background()

		// Generate token.
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		pay, err := paymentClient.Create(ctx, req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Capture payment.
		pay, err = paymentClient.Capture(ctx, pay.ID)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay != nil && pay.Status != "approved" {
			t.Error("payment should be approved, but is wasn't")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_capture_amount_payment", func(t *testing.T) {
		ctx := context.Background()

		// Generate token.
		token, err := test.GenerateCardToken(ctx, cardTokenClient)
		if err != nil {
			t.Error("fail to generate card token", err)
		}

		// Create payment.
		req := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}
		pay, err := paymentClient.Create(ctx, req)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Capture payment.
		pay, err = paymentClient.CaptureAmount(ctx, pay.ID, 100.1)
		if pay == nil {
			t.Error("payment can't be nil")
		}
		if pay != nil && pay.Status != "approved" {
			t.Error("payment should be approved, but is wasn't")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
