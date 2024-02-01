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

		client := payment.NewClient(c)

		dto := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		result, err := client.Create(context.Background(), dto)
		if result == nil {
			t.Error("result can't be nil")
		}
		if result.ID == 0 {
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

		dto := payment.SearchRequest{
			Filters: map[string]string{
				"external_reference": "abc_def_ghi_123_456123",
			},
		}

		client := payment.NewClient(c)
		result, err := client.Search(context.Background(), dto)
		if result == nil {
			t.Error("result can't be nil")
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

		client := payment.NewClient(c)
		dto := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		result, err := client.Create(context.Background(), dto)
		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		result, err = client.Get(context.Background(), result.ID)
		if result == nil {
			t.Error("result can't be nil")
		}
		if result.ID == 0 {
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

		client := payment.NewClient(c)
		dto := payment.Request{
			TransactionAmount: 105.1,
			PaymentMethodID:   "pix",
			Payer: &payment.PayerRequest{
				Email: fmt.Sprintf("gabs_%s@testuser.com", uuid.New()),
			},
		}

		result, err := client.Create(context.Background(), dto)
		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		result, err = client.Cancel(context.Background(), result.ID)
		if result == nil {
			t.Error("result can't be nil")
		}
		if result.ID == 0 {
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

		result, err := client.Create(context.Background(), dto)
		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}

		result, err = client.Capture(context.Background(), result.ID)
		if result == nil {
			t.Error("result can't be nil")
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

		client := payment.NewClient(c)
		result, err := client.CaptureAmount(context.Background(), 123, 100.1)
		if result == nil {
			t.Error("result can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
