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
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentClient := payment.NewClient(c)

		paymentRequest := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		payment, err := paymentClient.Create(context.Background(), paymentRequest)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if payment.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_search_payment", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentRequest := payment.SearchRequest{
			Filters: map[string]string{
				"external_reference": "abc_def_ghi_123_456123",
			},
		}

		paymentClient := payment.NewClient(c)
		paymentSearch, err := paymentClient.Search(context.Background(), paymentRequest)
		if paymentSearch == nil {
			t.Error("paymentSearch can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_get_payment", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentClient := payment.NewClient(c)
		paymentRequest := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		payment, err := paymentClient.Create(context.Background(), paymentRequest)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		payment, err = paymentClient.Get(context.Background(), payment.ID)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if payment.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_cancel_payment", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentClient := payment.NewClient(c)
		paymentRequest := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		payment, err := paymentClient.Create(context.Background(), paymentRequest)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		payment, err = paymentClient.Cancel(context.Background(), payment.ID)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if payment.ID == 0 {
			t.Error("id can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	// We should validate how to test capture and capture amount.
	t.Run("should_capture_payment", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentClient := payment.NewClient(c)

		// Create payment.
		paymentRequest := payment.Request{
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

		payment, err := paymentClient.Create(context.Background(), paymentRequest)
		if payment == nil {
			t.Error("payment can't be nil")
			return
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		payment, err = paymentClient.Capture(context.Background(), payment.ID)
		if payment == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})

	t.Run("should_capture_amount_payment", func(t *testing.T) {
		c, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		paymentClient := payment.NewClient(c)
		payment, err := paymentClient.CaptureAmount(context.Background(), 123, 100.1)
		if payment == nil {
			t.Error("payment can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
