package create

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

	var req = cardtoken.Request{
		SiteID:          "{{SITE_ID}}",
		CardNumber:      "{{CARD_NUMBER}}",
		ExpirationMonth: "11",
		ExpirationYear:  "2025",
		SecurityCode:    "123",
		Cardholder: &cardtoken.Cardholder{
			Identification: &cardtoken.Identification{
				Type:   "CPF",
				Number: "{{CPF_NUMBER}}",
			},
			Name: "{{PAYMENT_METHOD}}",
		},
	}

	result, err := client.Create(context.Background(), req)
	if err != nil {
		return
	}

	fmt.Println(result)
}
