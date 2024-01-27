package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/paymentmethod"
)

func TestPaymentMethod(t *testing.T) {
	t.Run("should_list_payment_methods", func(t *testing.T) {
		c, err := config.New(os.Getenv("at"))
		if err != nil {
			t.Fatal(err)
		}

		pmc := paymentmethod.NewClient(c)
		res, err := pmc.List(context.Background())

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
