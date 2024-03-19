package integration

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/refund"
	"github.com/mercadopago/sdk-go/test"
)

var (
	cfg             = test.Config()
	paymentClient   = payment.NewClient(cfg)
	cardTokenClient = cardtoken.NewClient(cfg)
	refundClient    = refund.NewClient(cfg)
)

func TestRefund(t *testing.T) {
	t.Run("should_create_refund", func(t *testing.T) {
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

		// Create refund.
		ref, err := refundClient.Create(ctx, pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_create_partial_refund", func(t *testing.T) {
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

		// Create partial refund.
		partialAmount := pay.TransactionAmount - 5.0

		ref, err := refundClient.CreatePartialRefund(ctx, pay.ID, partialAmount)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_refund", func(t *testing.T) {
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

		// Create refund.
		ref, err := refundClient.Create(ctx, pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Get refund.
		ref, err = refundClient.Get(ctx, pay.ID, ref.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if ref.ID == 0 {
			t.Error("id can't be nil")
		}
	})

	t.Run("should_list_refund", func(t *testing.T) {
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
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create partial refund.
		partialAmount := req.TransactionAmount - 5.0

		ref, err := refundClient.CreatePartialRefund(ctx, pay.ID, partialAmount)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create total refund.
		ref, err = refundClient.Create(ctx, pay.ID)
		if ref == nil {
			t.Error("refund can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// List refunds.
		refunds, err := refundClient.List(ctx, pay.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(refunds) != 2 {
			t.Error("size can't be different of 2")
		}
	})
}
