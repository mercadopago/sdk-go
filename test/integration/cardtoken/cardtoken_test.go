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

		client := cardtoken.NewClient(cfg)
		cardToken, err := client.Create(context.Background(), mockCardTokenRequest())

		if cardToken == nil {
			t.Error("cardToken can't be nil")
		}
		if err != nil {
			t.Errorf(err.Error())
		}
	})
}

func mockCardTokenRequest() cardtoken.Request {
	return cardtoken.Request{
		SiteID:          "Teste",
		CardNumber:      "5031433215406351",
		ExpirationMonth: "11",
		ExpirationYear:  "2025",
		SecurityCode:    "123",
		Cardholder: &cardtoken.CardholderRequest{
			Identification: &cardtoken.IdentificationRequest{
				Type:   "CPF",
				Number: "70383868084",
			},
			Name: "MASTER TEST",
		},
	}
}
