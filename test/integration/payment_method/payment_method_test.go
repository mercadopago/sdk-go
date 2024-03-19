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
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := paymentmethod.NewClient(cfg)
		resource, err := client.List(context.Background())

		if resource == nil {
			t.Error("resource can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
