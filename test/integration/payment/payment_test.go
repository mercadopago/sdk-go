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

		request := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
		}

		resource, err := paymentClient.Create(ctx, request)
		if resource == nil || resource.ID == 0 {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_payment", func(t *testing.T) {
		ctx := context.Background()

		request := payment.SearchRequest{
			Filters: map[string]string{
				"external_reference": "abc_def_ghi_123_456123",
			},
		}

		resource, err := paymentClient.Search(ctx, request)
		if resource == nil {
			t.Error("resource can't be nil")
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
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
		}

		resource, err := paymentClient.Create(ctx, paymentRequest)
		if resource == nil || resource.ID == 0 {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Get payment.
		resource, err = paymentClient.Get(ctx, resource.ID)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_cancel_payment", func(t *testing.T) {
		ctx := context.Background()

		// Create payment.
		request := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
		}

		resource, err := paymentClient.Create(ctx, request)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Cancel payment.
		resource, err = paymentClient.Cancel(ctx, resource.ID)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if resource != nil && resource.Status != "cancelled" {
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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		resource, err := paymentClient.Create(ctx, request)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Capture payment.
		resource, err = paymentClient.Capture(ctx, resource.ID)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if resource != nil && resource.Status != "approved" {
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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}
		resource, err := paymentClient.Create(ctx, request)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Capture payment.
		resource, err = paymentClient.CaptureAmount(ctx, resource.ID, 100.1)
		if resource == nil {
			t.Error("payment can't be nil")
		}
		if resource != nil && resource.Status != "approved" {
			t.Error("payment should be approved, but is wasn't")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
