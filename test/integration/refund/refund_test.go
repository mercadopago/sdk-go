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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		paymentResource, err := paymentClient.Create(ctx, request)
		if paymentResource == nil {
			t.Error("paymentResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Create refund.
		refundResource, err := refundClient.Create(ctx, paymentResource.ID)
		if refundResource == nil {
			t.Error("refundResource can't be nil")
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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		paymentResource, err := paymentClient.Create(ctx, request)
		if paymentResource == nil {
			t.Error("paymentResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Create partial refund.
		partialAmount := paymentResource.TransactionAmount - 5.0

		refundResource, err := refundClient.CreatePartialRefund(ctx, paymentResource.ID, partialAmount)
		if refundResource == nil {
			t.Error("refundResource can't be nil")
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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		paymentResource, err := paymentClient.Create(ctx, request)
		if paymentResource == nil {
			t.Error("paymentResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		// Create refund.
		refundResource, err := refundClient.Create(ctx, paymentResource.ID)
		if refundResource == nil {
			t.Error("refundResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Get refund.
		refundResource, err = refundClient.Get(ctx, paymentResource.ID, refundResource.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if refundResource.ID == 0 {
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
		request := payment.Request{
			TransactionAmount: 105.1,
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@meli.com", uuid.New()),
			},
			Token:        token,
			Installments: 1,
			Capture:      false,
		}

		paymentResource, err := paymentClient.Create(ctx, request)
		if paymentResource == nil {
			t.Error("paymentResource can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create partial refund.
		partialAmount := request.TransactionAmount - 5.0

		refundResource, err := refundClient.CreatePartialRefund(ctx, paymentResource.ID, partialAmount)
		if refundResource == nil {
			t.Error("refundResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create total refund.
		refundResource, err = refundClient.Create(ctx, paymentResource.ID)
		if refundResource == nil {
			t.Error("refundResource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// List refunds.
		resources, err := refundClient.List(ctx, paymentResource.ID)
		if err != nil {
			t.Errorf(err.Error())
		}

		if len(resources) != 2 {
			t.Error("size can't be different of 2")
		}
	})
}
