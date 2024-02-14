package cardtoken

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

		cardTokenClient := cardtoken.NewClient(cfg)
		cardToken, err := cardTokenClient.Create(context.Background(), cardtoken.MockCardTokenRequest())

		if cardToken == nil {
			t.Error("cardToken can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}
