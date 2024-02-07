package integration

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/payment/refund"
)

func TestRefund(t *testing.T) {
	// We need card token client to create a payment and then refund payment.
	t.Run("should_create_refund", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(c)

		// Create payment.
		dto := payment.Request{
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

		payment, err := client.Create(context.Background(), dto)
		if payment == nil {
			t.Error("result can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create refund.
		refundRequest := refund.Request{
			Amount: dto.TransactionAmount,
		}

		refundClient := refund.NewClient(c)
		refund, err := refundClient.Create(context.Background(), refundRequest, payment.ID)
		if refund == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// We need card token client to create a payment and then refund payment.
	t.Run("should_get_refund", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(c)

		// Create payment.
		dto := payment.Request{
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

		payment, err := client.Create(context.Background(), dto)
		if payment == nil {
			t.Error("result can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create refund.
		refundRequest := refund.Request{
			Amount: dto.TransactionAmount,
		}

		refundClient := refund.NewClient(c)
		refund, err := refundClient.Create(context.Background(), refundRequest, payment.ID)
		if refund == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		_, err = refundClient.Get(context.Background(), payment.ID, refund.ID)
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_list_refund", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := payment.NewClient(c)

		// Create payment.
		dto := payment.Request{
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

		payment, err := client.Create(context.Background(), dto)
		if payment == nil {
			t.Error("result can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		// Create refund.
		refundRequest := refund.Request{
			Amount: dto.TransactionAmount,
		}

		refundClient := refund.NewClient(c)
		refund, err := refundClient.Create(context.Background(), refundRequest, payment.ID)
		if refund == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		_, err = refundClient.List(context.Background(), payment.ID)
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
