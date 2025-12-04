package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/cardtoken"
	"github.com/mercadopago/sdk-go/pkg/config"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cardtoken.NewClient(cfg)

	request := cardtoken.Request{
		SiteID:          "{{SITE_ID}}",
		CardNumber:      "{{CARD_NUMBER}}",
		ExpirationMonth: "{{EXPIRATION_MONTH}}",
		ExpirationYear:  "{{EXPIRATION_YEAR}}",
		SecurityCode:    "{{SECURITY_CODE}}",
		Cardholder: &cardtoken.CardholderRequest{
			Identification: &cardtoken.IdentificationRequest{
				Type:   "CPF",
				Number: "{{NUMBER_CPF}}",
			},
			Name: "{{NAME}}",
		},
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		return
	}

	fmt.Println(resource)
}
