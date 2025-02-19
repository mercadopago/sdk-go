package test

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/cardtoken"
)

func GenerateCardToken(ctx context.Context, client cardtoken.Client) (string, error) {
	request := cardtoken.Request{
		Cardholder: &cardtoken.CardholderRequest{
			Identification: &cardtoken.IdentificationRequest{
				Number: "01234567890",
				Type:   "CPF",
			},
			Name: "APRO",
		},
		SiteID:          "MLB",
		CardNumber:      "5031433215406351",
		ExpirationYear:  "2030",
		ExpirationMonth: "11",
		SecurityCode:    "123",
	}

	resource, err := client.Create(ctx, request)
	if err != nil {
		return "", err
	}

	return resource.ID, nil
}
