package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func TestCardToken(t *testing.T) {
	t.Run("should_create_card_token", func(t *testing.T) {
		cfg, err := config.New(os.Getenv("ACCESS_TOKEN"))
		if err != nil {
			t.Fatal(err)
		}

		client := cardtoken.NewClient(cfg)
		res, err := client.Create(context.Background(), cardtoken.MockCardTokenRequest())

		if res == nil {
			t.Error("res can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
