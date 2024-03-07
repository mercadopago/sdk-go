package test

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/cardtoken"
)

func GenerateCardToken(ctx context.Context, client cardtoken.Client) (string, error) {
	req := cardtoken.Request{
		Cardholder: &cardtoken.Cardholder{
			Identification: &cardtoken.Identification{
				Number: "01234567890",
				Type:   "CPF",
			},
			Name: "APRO",
		},
		SiteID:          "MLB",
		CardNumber:      "5031433215406351",
		ExpirationYear:  "2025",
		ExpirationMonth: "11",
		SecurityCode:    "123",
	}

	result, err := client.Create(context.Background(), req)
	if err != nil {
		return "", err
	}

	return result.ID, nil
}
