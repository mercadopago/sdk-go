package main

import (
	"context"
	"fmt"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

func main() {
	cfg, err := config.New("APP_USR-3031066189562927-042217-7204e46c8ca7b09cee4327a6e38f8c0d-831921084")
	if err != nil {
		fmt.Println(err)
		return
	}

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				ID:          "123",
				Title:       "Title",
				UnitPrice:   75,
				Quantity:    1,
				Description: "Description",
			},
		},
		PaymentMethods: &preference.PaymentMethodsRequest{
			DefaultPaymentMethodID: "visa",
		},
		Payer: &preference.PayerRequest{
			Name:    "tyuio",
			Surname: "gfhjkl",
			Email:   "test_user_61213998@testuser.com",
			Identification: &preference.IdentificationRequest{
				Type:   "CPF",
				Number: "19119119100",
			},
			Address: &preference.AddressRequest{
				ZipCode:      "88054000",
				StreetName:   "Rua PREFERENCE",
				StreetNumber: "12345",
			},
		},
	}

	client := preference.NewClient(cfg)
	resource, err := client.Create(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
